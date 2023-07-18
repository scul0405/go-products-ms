package grpc

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	incomingMessage = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_incoming_grpc_messages_total",
		Help: "The total number of incoming gRPC messages",
	})

	successMessage = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_success_incoming_grpc_messages_total",
		Help: "The total number of success incoming gRPC messages",
	})

	errorMessage = promauto.NewCounter(prometheus.CounterOpts{
		Name: "products_error_incoming_grpc_message_total",
		Help: "The total number of error incoming gRPC messages",
	})
)
