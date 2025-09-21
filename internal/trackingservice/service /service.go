package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"wheres-my-pizza/internal/trackingservice/db"
	"wheres-my-pizza/pkg/config"
	"wheres-my-pizza/pkg/logger"
)

type TrackingService struct {
	DB     *db.TrackingDB
	Logger *logger.Logger
}

func NewTrackingService(db *db.TrackingDB, logger *logger.Logger) *TrackingService {
	return &TrackingService{
		DB:     db,
		Logger: logger,
	}
}

// Handler to get the current status of an order
func (s *TrackingService) GetOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Extract order_number from the URL path
	orderNumber := extractOrderNumber(r.URL.Path)

	// Query the database for the order
	order, err := s.DB.GetOrderByNumber(r.Context(), orderNumber)
	if err != nil {
		if err == db.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		s.Logger.Error("db_query_failed", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Construct response body
	status := struct {
		OrderNumber         string    `json:"order_number"`
		CurrentStatus       string    `json:"current_status"`
		UpdatedAt           time.Time `json:"updated_at"`
		EstimatedCompletion time.Time `json:"estimated_completion"`
		ProcessedBy         string    `json:"processed_by"`
	}{
		OrderNumber:         order.Number,
		CurrentStatus:       order.Status,
		UpdatedAt:           order.UpdatedAt,
		EstimatedCompletion: order.CompletedAt, // Assuming this is the estimated completion time
		ProcessedBy:         order.ProcessedBy,
	}

	// Return the response in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

// Handler to get the history of an order
func (s *TrackingService) GetOrderHistory(w http.ResponseWriter, r *http.Request) {
	// Extract order_number from the URL path
	orderNumber := extractOrderNumber(r.URL.Path)

	// Query the database for the order
	order, err := s.DB.GetOrderByNumber(r.Context(), orderNumber)
	if err != nil {
		if err == db.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		s.Logger.Error("db_query_failed", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get the order status history
	history, err := s.DB.GetOrderStatusHistory(r.Context(), order.ID)
	if err != nil {
		s.Logger.Error("db_query_failed", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Construct response body
	var historyResponse []struct {
		Status    string    `json:"status"`
		Timestamp time.Time `json:"timestamp"`
		ChangedBy string    `json:"changed_by"`
	}
	for _, entry := range history {
		historyResponse = append(historyResponse, struct {
			Status    string    `json:"status"`
			Timestamp time.Time `json:"timestamp"`
			ChangedBy string    `json:"changed_by"`
		}{
			Status:    entry.Status,
			Timestamp: entry.CreatedAt,
			ChangedBy: entry.ChangedBy,
		})
	}

	// Return the response in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historyResponse)
}

// Handler to get the status of all kitchen workers
func (s *TrackingService) GetWorkersStatus(w http.ResponseWriter, r *http.Request) {
	// Query the database for the workers
	workers, err := s.DB.GetAllWorkers(r.Context())
	if err != nil {
		s.Logger.Error("db_query_failed", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Construct response body
	var workersResponse []struct {
		WorkerName      string    `json:"worker_name"`
		Status          string    `json:"status"`
		OrdersProcessed int       `json:"orders_processed"`
		LastSeen        time.Time `json:"last_seen"`
	}
	for _, worker := range workers {
		// Logic to determine if a worker is offline based on last_seen
		status := "online"
		if time.Since(worker.LastSeen) > time.Hour*2 { // Example: offline if last seen more than 2 hours ago
			status = "offline"
		}

		workersResponse = append(workersResponse, struct {
			WorkerName      string    `json:"worker_name"`
			Status          string    `json:"status"`
			OrdersProcessed int       `json:"orders_processed"`
			LastSeen        time.Time `json:"last_seen"`
		}{
			WorkerName:      worker.Name,
			Status:          status,
			OrdersProcessed: worker.OrdersProcessed,
			LastSeen:        worker.LastSeen,
		})
	}

	// Return the response in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workersResponse)
}

// Utility function to extract order_number from the URL
func extractOrderNumber(path string) string {
	// Assumes the path is of the form /orders/{order_number}/status or /orders/{order_number}/history
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return ""
	}
	return parts[2] // The order number will be the third part
}

func StartTrackingService(cfg *config.Config, logger *logger.Logger) error {
	// Initialize DB connection pool and the TrackingService
	dbPool, err := db.NewDBPool(cfg.DatabaseURL) // Implement NewDBPool to initialize DB connection
	if err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	trackingDB := db.NewTrackingDB(dbPool, logger)
	service := NewTrackingService(trackingDB, logger)

	// Set up the default multiplexer
	mux := http.NewServeMux()

	// Define the routes
	mux.HandleFunc("/orders/", service.GetOrderStatus)  // Will handle /orders/{order_number}/status
	mux.HandleFunc("/orders/", service.GetOrderHistory) // Will handle /orders/{order_number}/history
	mux.HandleFunc("/workers/status", service.GetWorkersStatus)

	// Start the HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	logger.Info("service_started", nil)
	return server.ListenAndServe()
}
