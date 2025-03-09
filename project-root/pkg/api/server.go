package api

import (
	"log"
	"net/http"

	"project-root/pkg/api/auth"
)

func StartServer(addr, certFile, keyFile string) {
	mux := http.NewServeMux()

	// API endpoints with authentication middleware.
	mux.HandleFunc("/api/report", auth.AuthMiddleware(reportHandler))
	mux.HandleFunc("/api/status", auth.AuthMiddleware(statusHandler))

	log.Printf("Starting API server on %s", addr)
	err := http.ListenAndServeTLS(addr, certFile, keyFile, mux)
	if err != nil {
		log.Fatalf("API server failed: %v", err)
	}
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse agent report (likely JSON) and update the database.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Report received"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Simple status check for agents.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Agent is online"))
}
