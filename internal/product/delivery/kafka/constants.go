package kafka

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

const (
	minBytes               = 10e3
	maxBytes               = 10e6
	queueCapacity          = 100
	heartbeatInterval      = 3 * time.Second
	commitInterval         = 0
	partitionWatchInterval = 5 * time.Second
	maxAttempts            = 3
	dialTimeout            = 3 * time.Minute

	writerReadTimeout  = 10 * time.Second
	writerWriteTimeout = 10 * time.Second

	createProductTopic   = "create-product"
	createProductWorkers = 16
	updateProductTopic   = "update-product"
	updateProductWorkers = 16

	deadLetterQueueTopic = "dead-letter-queue"

	productsGroupID = "products_group"
)

var (
	incomingMessage = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_incoming_kafka_messages_total",
		Help: "The total number of incoming Kafka messages",
	})

	successMessage = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_success_incoming_kafka_messages_total",
		Help: "The total number of success incoming Kafka messages",
	})

	errorMessage = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_error_incoming_kafka_message_total",
		Help: "The total number of error incoming Kafka messages",
	})
)
