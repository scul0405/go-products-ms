package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	AppVersion string
	Server     Server
	Logger     Logger
	MongoDB    MongoDB
	Metrics    Metrics
	Jaeger     Jaeger
	Kafka      Kafka
}

// Server config
type Server struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

// Logger config
type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Metrics config
type Metrics struct {
	Port        string
	URL         string
	ServiceName string
}

// Jaeger config
type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

// MongoDB config
type MongoDB struct {
	URI string
	DB  string
}

// Kafka config
type Kafka struct {
	Brokers []string
}

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
