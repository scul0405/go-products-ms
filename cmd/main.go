package main

import (
	"Go-ProductMS/config"
	"Go-ProductMS/internal/server"
	"Go-ProductMS/pkg/logger"
	"Go-ProductMS/pkg/mongodb"
	"context"
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

	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
