package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Обработка сигналов для graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("Получен сигнал завершения, закрываю status-service...")
		cancel()
	}()

	// Получаем URL NATS из переменной окружения или используем значение по умолчанию
	natsURL := getEnv("NATS_URL", "nats://localhost:4222")
	log.Printf("Подключение к NATS серверу: %s", natsURL)

	// Подключаемся к NATS серверу
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
	defer nc.Close()

	log.Println("Успешно подключен к NATS серверу")

	// Запускаем периодический запрос статуса
	go checkStatus(ctx, nc)

	// Ожидаем завершения контекста
	<-ctx.Done()
	log.Println("Status service завершил работу")
}

func checkStatus(ctx context.Context, nc *nats.Conn) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Отправляем запрос на получение статуса
			msg, err := nc.Request("jobs.status", []byte("status"), 2*time.Second)
			if err != nil {
				log.Printf("Ошибка при запросе статуса: %v", err)
			} else {
				log.Printf("Получен статус: %s", string(msg.Data))
			}
		}
	}
}

// Вспомогательная функция для получения переменных окружения
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
