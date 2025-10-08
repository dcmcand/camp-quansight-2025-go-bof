package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dcmcand/camp-quansight-2025-go-bof/pkg/even"
	"github.com/google/uuid"
)

type NumberRequest struct {
	Number int `json:"number"`
}

type EvenResponse struct {
	IsEven bool `json:"is_even"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

type contextKey string

const requestIDKey contextKey = "requestID"

// requestIDMiddleware generates a unique request ID and adds it to the context
func requestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// getRequestID retrieves the request ID from context
func getRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return "unknown"
}

func isEvenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := getRequestID(ctx)

	slog.Info("Request received", "request_id", requestID, "method", r.Method, "path", r.URL.Path, "remote_addr", r.RemoteAddr)

	if r.Method != http.MethodPost {
		slog.Warn("Method not allowed", "request_id", requestID, "method", r.Method, "path", r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req NumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Invalid request body", "request_id", requestID, "error", err, "path", r.URL.Path)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	isEven, err := even.IsEven(ctx, req.Number)
	if err != nil {
		slog.Error("Error processing IsEven request", "request_id", requestID, "error", err, "number", req.Number)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	response := EvenResponse{IsEven: isEven}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	slog.Info("Request completed successfully", "request_id", requestID, "number", req.Number, "is_even", isEven)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := getRequestID(ctx)

	slog.Info("Health check received", "request_id", requestID, "method", r.Method, "path", r.URL.Path, "remote_addr", r.RemoteAddr)

	if r.Method != http.MethodGet {
		slog.Warn("Method not allowed on health endpoint", "request_id", requestID, "method", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	slog.Debug("Health check completed successfully", "request_id", requestID)
}

func ListenaAndServe(port string) error {
	http.HandleFunc("/", requestIDMiddleware(isEvenHandler))
	http.HandleFunc("/health", requestIDMiddleware(healthHandler))

	addr := fmt.Sprintf(":%s", port)
	slog.Info("Starting server", "port", port, "address", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("Server failed to start", "error", err, "port", port)
		return err
	}
	return nil
}
