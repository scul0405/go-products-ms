package kafka

import (
	"Go-ProductMS/config"
	"Go-ProductMS/internal/models"
	"Go-ProductMS/internal/product/usecase"
	"Go-ProductMS/pkg/logger"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"sync"
)

type ProductsConsumerGroup struct {
	Brokers    []string
	GroupID    string
	log        logger.Logger
	cfg        *config.Config
	productsUC usecase.ProductsUseCase
}

func NewProductsConsumerGroup(
	brokers []string,
	groupID string,
	log logger.Logger,
	cfg *config.Config,
	productsUC usecase.ProductsUseCase) *ProductsConsumerGroup {
	return &ProductsConsumerGroup{
		Brokers:    brokers,
		GroupID:    groupID,
		log:        log,
		cfg:        cfg,
		productsUC: productsUC,
	}
}

func (pcg *ProductsConsumerGroup) getNewKafkaReader(kafkaURL []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		Topic:                  topic,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		Logger:                 kafka.LoggerFunc(pcg.log.Debugf),
		ErrorLogger:            kafka.LoggerFunc(pcg.log.Errorf),
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}

func (pcg *ProductsConsumerGroup) getNewKafkaWriter(topic string) *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(pcg.Brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: -1,
		MaxAttempts:  3,
		Logger:       kafka.LoggerFunc(pcg.log.Infof),
		ErrorLogger:  kafka.LoggerFunc(pcg.log.Errorf),
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
	}
	return w
}

func (pcg *ProductsConsumerGroup) consumeCreateProduct(ctx context.Context, cancel context.CancelFunc, groupID string, topic string, workersNum int) {
	defer cancel()

	reader := pcg.getNewKafkaReader(pcg.Brokers, topic, groupID)
	defer func() {
		if err := reader.Close(); err != nil {
			pcg.log.Errorf("Error while closing reader: %s", err)
			cancel()
		}
	}()

	writer := pcg.getNewKafkaWriter(deadLetterQueueTopic)
	defer func() {
		if err := writer.Close(); err != nil {
			pcg.log.Errorf("w.Close", err)
			cancel()
		}
	}()

	pcg.log.Infof("Starting consumer group: %v", reader.Config().GroupID)

	wg := &sync.WaitGroup{}

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go pcg.createProductWorker(ctx, cancel, reader, writer, wg, i)
	}
	wg.Wait()

}

func (pcg *ProductsConsumerGroup) consumeUpdateProduct(ctx context.Context, cancel context.CancelFunc, groupID string, topic string, workersNum int) {
	reader := pcg.getNewKafkaReader(pcg.Brokers, topic, groupID)
	defer cancel()
	defer func() {
		if err := reader.Close(); err != nil {
			pcg.log.Errorf("Error while closing reader: %s", err)
			cancel()
		}
	}()

	writer := pcg.getNewKafkaWriter(deadLetterQueueTopic)
	defer func() {
		if err := writer.Close(); err != nil {
			pcg.log.Errorf("w.Close", err)
			cancel()
		}
	}()

	pcg.log.Infof("Starting consumer group: %v", reader.Config().GroupID)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go pcg.updateProductWorker(ctx, cancel, reader, writer, wg, i)
	}
	wg.Wait()
}

func (pcg *ProductsConsumerGroup) publishErrorMessage(ctx context.Context, w *kafka.Writer, m kafka.Message, err error) error {
	errMsg := &models.ErrorMessage{
		Offset:    m.Offset,
		Error:     err.Error(),
		Time:      m.Time.UTC(),
		Partition: m.Partition,
		Topic:     m.Topic,
	}

	errMsgBytes, err := json.Marshal(errMsg)
	if err != nil {
		return err
	}

	return w.WriteMessages(ctx, kafka.Message{
		Value: errMsgBytes,
	})
}

func (pcg *ProductsConsumerGroup) RunConsumers(ctx context.Context, cancel context.CancelFunc) {
	go pcg.consumeCreateProduct(ctx, cancel, productsGroupID, createProductTopic, createProductWorkers)
	go pcg.consumeUpdateProduct(ctx, cancel, productsGroupID, updateProductTopic, updateProductWorkers)
}
