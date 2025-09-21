package orderservice

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wheres-my-pizza/cmd/orderservice/server"
	"wheres-my-pizza/pkg/config"
	"wheres-my-pizza/pkg/logger"
)

func Main() {
	port := flag.Int("port", 3000, "The HTTP port for the API.")
	maxConcurrent := flag.Int("max-concurrent", 50, "Maximum number of concurrent orders to process.")
	flag.Parse()

	logger := logger.NewLogger("orderservice")

	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Error("startup-001", "config_load_failed", "Failed to load configuration", err)
		log.Fatal(err)
	}

	srv := server.NewServer(*port, *maxConcurrent, cfg, logger)

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Error("startup-001", "server_start_failed", "Failed to start server", err)
			log.Fatal(err)
		}
	}()

	// Graceful shutfown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutdown", "graceful_shutdown", "Shutting down server...", map[string]interface{}{})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("shutdown", "shutdown_failed", "Server forced to shutdown", err)
		log.Fatal(err)
	}

	logger.Info("shutdown", "service_stopped", "Server exiting", map[string]interface{}{})

}
