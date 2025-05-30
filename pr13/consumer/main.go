package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	tasksProcessed int
	mu             sync.Mutex
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
		log.Println("Получен сигнал завершения, закрываю consumer...")
		cancel()
	}()

	// Получаем параметры из переменных окружения
	natsURL := getEnv("NATS_URL", "nats://localhost:4222")
	consumerID := getEnv("CONSUMER_ID", "consumer-default")

	log.Printf("Consumer %s: подключение к NATS серверу: %s", consumerID, natsURL)

	// Подключаемся к NATS серверу
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
	defer nc.Close()

	log.Printf("Consumer %s: успешно подключен к NATS серверу", consumerID)

	// Инициализируем счетчик обработанных задач
	tasksProcessed = 0

	// Подписываемся на задачи с использованием Queue Group
	sub, err := nc.QueueSubscribe("jobs.create", "workers", func(msg *nats.Msg) {
		task := string(msg.Data)
		log.Printf("Consumer %s: получена задача: %s", consumerID, task)

		// Имитируем обработку задачи
		processingTime := 500 + rand.Intn(1000) // от 500 до 1500 мс
		time.Sleep(time.Duration(processingTime) * time.Millisecond)

		mu.Lock()
		tasksProcessed++
		currentTasks := tasksProcessed
		mu.Unlock()

		log.Printf("Consumer %s: задача %s обработана за %dмс (всего: %d)",
			consumerID, task, processingTime, currentTasks)
	})
	if err != nil {
		log.Fatalf("Ошибка подписки на тему: %v", err)
	}

	// Подписываемся на запросы статуса
	statusSub, err := nc.QueueSubscribe("jobs.status", "status-workers", func(msg *nats.Msg) {
		mu.Lock()
		currentTasks := tasksProcessed
		mu.Unlock()

		response := []byte(consumerID + ": обработано задач: " + string(rune(currentTasks)))
		if err := msg.Respond(response); err != nil {
			log.Printf("Ошибка при отправке ответа: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Ошибка подписки на тему статуса: %v", err)
	}

	// Ожидаем завершения контекста
	<-ctx.Done()
	log.Printf("Consumer %s: отписка от темы и завершение работы", consumerID)
	sub.Unsubscribe()
	statusSub.Unsubscribe()
}

// Вспомогательная функция для получения переменных окружения
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
