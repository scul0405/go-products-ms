package mongodb

import (
	"Go-ProductMS/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	connectionTimeout = 30 * time.Second
	minConnIdleTime   = 3 * time.Minute
	minPoolSize       = 20
	maxPoolSize       = 300
)

func NewMongoDBConnection(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(cfg.MongoDB.URI).
		SetConnectTimeout(connectionTimeout).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetMaxConnIdleTime(minConnIdleTime)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err != client.Ping(ctx, nil) {
		return nil, err
	}

	return client, nil
}
