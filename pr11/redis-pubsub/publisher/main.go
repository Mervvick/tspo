package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Бесконечный цикл публикации сообщений
	counter := 1
	for {
		message := fmt.Sprintf("Hello #%d - %s", counter, time.Now().Format(time.RFC3339))
		err := rdb.Publish(ctx, "notifications", message).Err()
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			log.Printf("Published: %s", message)
		}

		counter++
		time.Sleep(2 * time.Second)
	}
}
