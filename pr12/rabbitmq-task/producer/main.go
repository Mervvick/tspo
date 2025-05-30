package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	amqp "github.com/rabbitmq/amqp091-go"
)

type EmailTask struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Подключение к RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Создаем mutex для синхронизации доступа к каналу
	var channelMutex sync.Mutex

	// Настройка роутера
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/send-email", func(w http.ResponseWriter, r *http.Request) {
		var task EmailTask
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Проверка обязательных полей
		if task.To == "" || task.Subject == "" {
			http.Error(w, "To and Subject fields are required", http.StatusBadRequest)
			return
		}

		// Сериализация задачи
		body, err := json.Marshal(task)
		if err != nil {
			http.Error(w, "Error serializing task", http.StatusInternalServerError)
			return
		}

		// Блокируем доступ к каналу для текущего запроса
		channelMutex.Lock()
		defer channelMutex.Unlock()

		// Создаем новый канал для каждого запроса
		ch, err := conn.Channel()
		if err != nil {
			http.Error(w, "Failed to open a channel", http.StatusInternalServerError)
			log.Printf("Failed to open a channel: %s", err)
			return
		}
		defer ch.Close()

		// Включаем publisher confirms
		err = ch.Confirm(false) // noWait = false
		if err != nil {
			http.Error(w, "Failed to enable publisher confirms", http.StatusInternalServerError)
			log.Printf("Failed to enable publisher confirms: %s", err)
			return
		}

		// Создание очереди с настройками долговечности
		q, err := ch.QueueDeclare(
			"email_tasks", // name
			true,          // durable - очередь сохранится после перезапуска сервера
			false,         // delete when unused
			false,         // exclusive
			false,         // no-wait
			nil,           // arguments
		)
		if err != nil {
			http.Error(w, "Failed to declare a queue", http.StatusInternalServerError)
			log.Printf("Failed to declare a queue: %s", err)
			return
		}

		// Канал для подтверждений
		confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

		// Установка таймаута публикации
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Публикация сообщения с настройками долговечности
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			true,   // mandatory - сообщение вернется, если его некуда маршрутизировать
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, // сообщение сохранится при перезапуске сервера
				ContentType:  "application/json",
				Body:         body,
			})
		if err != nil {
			http.Error(w, "Failed to publish message", http.StatusInternalServerError)
			log.Printf("Failed to publish message: %s", err)
			return
		}

		// Ожидаем подтверждения от сервера с таймаутом
		select {
		case confirm := <-confirms:
			if confirm.Ack {
				log.Printf("Message confirmed by server [%d]", confirm.DeliveryTag)
				fmt.Fprintf(w, "Email task queued successfully")
			} else {
				log.Printf("Message not confirmed by server [%d]", confirm.DeliveryTag)
				http.Error(w, "Failed to queue message - no confirmation", http.StatusInternalServerError)
			}
		case <-time.After(2 * time.Second):
			log.Printf("No confirmation received from server (timeout)")
			http.Error(w, "Failed to queue message - confirmation timeout", http.StatusInternalServerError)
		}
	})

	log.Println("Email task producer is running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
