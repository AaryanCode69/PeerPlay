package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlerHealth)
	mux.HandleFunc("POST /api/echo", handleRequestBody)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           middlewareLogRequest(mux),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
