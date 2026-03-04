package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Response struct {
	Reason string `json:"reason"`
}

func noHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	reason := reasons[rand.Intn(len(reasons))]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Reason: reason})
}