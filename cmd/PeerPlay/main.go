package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"PeerPlay/internal/config"
	"PeerPlay/internal/handler"
	"PeerPlay/internal/middleware"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("invalid config : %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handler.HandlerHealth)
	mux.HandleFunc("POST /api/echo", handler.HandleRequestBody)

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           middleware.MiddlewareLogRequest(mux),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
