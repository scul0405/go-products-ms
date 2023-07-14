package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppVersion string
	Server     Server
	Logger     Logger
}

type Server struct {
	Port        string
	Development bool
}

type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
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
