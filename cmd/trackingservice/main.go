package trackingservice

import (
	"log"
	"wheres-my-pizza/internal/trackingservice/service"
	"wheres-my-pizza/pkg/config"
	"wheres-my-pizza/pkg/logger"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	// Инициализация логгера
	log := logger.NewLogger()

	// Запуск сервиса отслеживания
	if err := service.StartTrackingService(cfg, log); err != nil {
		log.Fatalf("Ошибка при запуске сервиса: %v", err)
	}
}
