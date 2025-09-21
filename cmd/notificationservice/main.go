package notificationservice

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wheres-my-pizza/internal/notificationservice/subscriber"
	"wheres-my-pizza/pkg/config"
	"wheres-my-pizza/pkg/logger"
)

func Main() {
	logger := logger.NewLogger("notification-subscriber")
	logger.Info("startup-001", "service_started", "Notification Subscriber starting", map[string]interface{}{})

	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Error("startup", "config_load_failed", "Failed to load configuration", err)
		log.Fatal(err)
	}

	nsub := subscriber.NewNotificationSubscriber(cfg, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := nsub.Start(ctx); err != nil {
			logger.Error("startup", "subscribe_start_failed", "Failed to start subscriber", err)
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutdown", "graceful_shutdown", "Shutting down subscriber...", map[string]interface{}{})
	cancel()
	nsub.Stop()

	logger.Info("shutdown", "service_stopped", "Subscriber exiting", map[string]interface{}{})

}
