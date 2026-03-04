package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	loadReasons()
	go cleanupVisitors()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", loggingMiddleware(indexHandler))
	http.HandleFunc("/no", loggingMiddleware(rateLimitMiddleware(noHandler)))
	http.HandleFunc("/health", loggingMiddleware(healthHandler))

	log.Printf("server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}