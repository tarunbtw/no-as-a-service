package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	loadReasons()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "no-as-a-service is running")
	})
	http.HandleFunc("/no", noHandler)

	log.Printf("server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}