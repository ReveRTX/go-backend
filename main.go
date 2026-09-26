package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
		return
	}

	now := time.Now().In(loc)

	response := HealthResponse{
		Status:    "healthy",
		Message:   "Backend is running",
		Timestamp: now.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode health response: %v", err)
		return
	}

	log.Printf("Health check: %s %s - status=healthy timestamp=%s",
		r.Method,
		r.URL.Path,
		now.Format(time.RFC3339),
	)
}

func main() {
	http.HandleFunc("/health", healthHandler)

	port := ":8000"

	log.Printf("Backend starting on http://localhost%s", port)
	log.Printf("Health endpoint: http://localhost%s/health", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
