package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Reason string `json:"reason"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Reasons int    `json:"reasons_loaded"`
	Uptime  string `json:"uptime"`
}

var startTime = time.Now()

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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status:  "ok",
		Reasons: len(reasons),
		Uptime:  time.Since(startTime).Round(time.Second).String(),
	})
}