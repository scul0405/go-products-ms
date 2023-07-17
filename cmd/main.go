package main

import (
	"Go-ProductMS/config"
	"Go-ProductMS/internal/server"
	"Go-ProductMS/pkg/jaeger"
	"Go-ProductMS/pkg/kafka"
	"Go-ProductMS/pkg/logger"
	"Go-ProductMS/pkg/mongodb"
	"context"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main() {
	log.Println("Starting products service...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("unable to parse config, %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting products service...")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, DevelopmentMode: %s",
		cfg.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Development,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.AppVersion)

	tracer, closer, err := jaeger.InitJaeger(cfg)
	if err != nil {
		appLogger.Fatal("cannot create tracer", err)
	}
	appLogger.Info("Jaeger connected")

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	appLogger.Info("Opentracing connected")

	mongoDBConn, err := mongodb.NewMongoDBConnection(ctx, cfg)
	if err != nil {
		appLogger.Fatalf("unable to connect to MongoDB, %v", err)
	}

	defer func() {
		if err := mongoDBConn.Disconnect(ctx); err != nil {
			appLogger.Fatalf("unable to disconnect from MongoDB, %v", err)
		}
	}()

	appLogger.Info("Success connected to MongoDB")

	conn, err := kafka.NewKafkaConn(cfg)
	if err != nil {
		appLogger.Fatal("NewKafkaConn", err)
	}
	defer conn.Close()
	brokers, err := conn.Brokers()
	if err != nil {
		appLogger.Fatal("conn.Brokers", err)
	}
	appLogger.Infof("Kafka connected: %v", brokers)

	s := server.NewServer(appLogger, cfg, mongoDBConn, tracer)
	appLogger.Fatal(s.Run())
}
