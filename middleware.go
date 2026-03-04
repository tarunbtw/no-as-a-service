package main

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next(wrapped, r)

		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", wrapped.status,
			"latency", time.Since(start).String(),
			"ip", r.RemoteAddr,
		)
	}
}