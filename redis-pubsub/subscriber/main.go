package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Подключение к Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // нет пароля
		DB:       0,  // используем базу данных по умолчанию
	})

	// Проверка соединения
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")

	// Подписка на канал notifications
	pubsub := rdb.Subscribe(ctx, "notifications")
	defer pubsub.Close()

	// Получаем канал для сообщений
	channel := pubsub.Channel()

	log.Println("Subscribed to notifications channel. Waiting for messages...")

	// Читаем сообщения из канала
	for msg := range channel {
		log.Printf("Received message from channel %s: %s", msg.Channel, msg.Payload)
	}
}
