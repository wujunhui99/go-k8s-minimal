package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	Port    string `json:"port"`
}

func loadConfig() Config {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "go-k8s-minimal"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		AppName: appName,
		Port:    port,
	}
}

func main() {
	cfg := loadConfig()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong"))
	})

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(cfg)
	})

	addr := ":" + cfg.Port
	log.Printf("server starting at %s, app=%s", addr, cfg.AppName)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}