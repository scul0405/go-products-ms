package kafka

import (
	"Go-ProductMS/internal/models"
	"context"
	"encoding/json"
	"github.com/avast/retry-go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/segmentio/kafka-go"
	"sync"
	"time"
)

const (
	retryAttempts = 1
	retryDelay    = 1 * time.Second
)

func (pcg *ProductsConsumerGroup) createProductWorker(
	ctx context.Context,
	cancel context.CancelFunc,
	r *kafka.Reader,
	w *kafka.Writer,
	wg *sync.WaitGroup,
	workerID int) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ProductsConsumerGroup.createProductWorker")
	defer span.Finish()

	span.LogFields(log.String("ConsumerGroup", r.Config().GroupID))

	defer wg.Done()
	defer cancel()

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			pcg.log.Errorf("FetchMessage", err)
			return
		}

		pcg.log.Infof(
			"WORKER: %v, message at topic/partition/offset %v/%v/%v: %s = %s\n",
			workerID,
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)
		incomingMessage.Inc()

		var prod models.Product
		if err := json.Unmarshal(m.Value, &prod); err != nil {
			errorMessage.Inc()
			pcg.log.Errorf("json.Unmarshal", err)
			continue
		}

		if err := pcg.validate.StructCtx(ctx, prod); err != nil {
			errorMessage.Inc()
			pcg.log.Errorf("validate.StructCtx", err)
		}

		if err := retry.Do(func() error {
			created, err := pcg.productsUC.Create(ctx, &prod)
			if err != nil {
				return err
			}
			pcg.log.Infof("created product: %v", created)
			return nil
		},
			retry.Attempts(retryAttempts),
			retry.Delay(retryDelay),
			retry.Context(ctx),
		); err != nil {
			errorMessage.Inc()

			if err := pcg.publishErrorMessage(ctx, w, m, err); err != nil {
				pcg.log.Errorf("publishErrorMessage", err)
				continue
			}
			pcg.log.Errorf("productsUC.Create.publishErrorMessage", err)
			continue
		}

		if err := r.CommitMessages(ctx, m); err != nil {
			errorMessage.Inc()
			pcg.log.Errorf("CommitMessages", err)
			continue
		}

		successMessage.Inc()
	}
}

func (pcg *ProductsConsumerGroup) updateProductWorker(
	ctx context.Context,
	cancel context.CancelFunc,
	r *kafka.Reader,
	w *kafka.Writer,
	wg *sync.WaitGroup,
	workerID int) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ProductsConsumerGroup.updateProductWorker")
	defer span.Finish()

	span.LogFields(log.String("ConsumerGroup", r.Config().GroupID))

	defer wg.Done()
	defer cancel()

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			pcg.log.Errorf("FetchMessage", err)
			return
		}

		pcg.log.Infof(
			"WORKER: %v, message at topic/partition/offset %v/%v/%v: %s = %s\n",
			workerID,
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)
		incomingMessage.Inc()

		var prod models.Product
		if err := json.Unmarshal(m.Value, &prod); err != nil {
			errorMessage.Inc()
			pcg.log.Errorf("json.Unmarshal", err)
			continue
		}

		if err := pcg.validate.StructCtx(ctx, prod); err != nil {
			errorMessage.Inc()
			pcg.log.Errorf("validate.StructCtx", err)
		}

		if err := retry.Do(func() error {
			updated, err := pcg.productsUC.Update(ctx, &prod)
			if err != nil {
				return err
			}
			pcg.log.Debugf("updated product: %v", updated)
			return nil
		},
			retry.Attempts(retryAttempts),
			retry.Delay(retryDelay),
			retry.Context(ctx),
		); err != nil {
			errorMessage.Inc()

			if err := pcg.publishErrorMessage(ctx, w, m, err); err != nil {
				pcg.log.Errorf("publishErrorMessage", err)
				continue
			}
			pcg.log.Errorf("productsUC.Update.publishErrorMessage", err)
			continue
		}

		if err := r.CommitMessages(ctx, m); err != nil {
			errorMessage.Inc()
			pcg.log.Errorf("CommitMessages", err)
			continue
		}

		successMessage.Inc()
	}
}
