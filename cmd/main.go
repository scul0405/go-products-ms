package main

import (
	"Go-ProductMS/config"
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
