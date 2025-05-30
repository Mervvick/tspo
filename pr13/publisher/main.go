package main

import (
	"context"
	"fmt"
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
		log.Println("Получен сигнал завершения, закрываю publisher...")
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

	// Отправляем задачи
	go publishTasks(ctx, nc)

	// Ожидаем завершения контекста
	<-ctx.Done()
	log.Println("Publisher завершил работу")
}

func publishTasks(ctx context.Context, nc *nats.Conn) {
	taskCount := 0
	for {
		select {
		case <-ctx.Done():
			return
		default:
			taskCount++
			taskMsg := fmt.Sprintf("Task %d", taskCount)
			err := nc.Publish("jobs.create", []byte(taskMsg))
			if err != nil {
				log.Printf("Ошибка при отправке задачи: %v", err)
			} else {
				log.Printf("Отправлена задача: %s", taskMsg)
			}
			time.Sleep(1 * time.Second) // Отправляем задачу каждую секунду
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
