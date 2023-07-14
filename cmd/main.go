package main

import (
	"Go-ProductMS/config"
	"Go-ProductMS/pkg/logger"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting products service...")

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

	http.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %v", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "ok"}); err != nil {
			log.Printf("unable to encode response, %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Fatal(http.ListenAndServe(cfg.Server.Port, nil))
	})
}
