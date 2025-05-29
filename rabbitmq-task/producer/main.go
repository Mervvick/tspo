package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Создание очереди
	q, err := ch.QueueDeclare(
		"email_tasks", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

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

		// Установка таймаута публикации
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Публикация сообщения
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         body,
			})
		failOnError(err, "Failed to publish a message")

		log.Printf("Sent email task: %+v", task)
		fmt.Fprintf(w, "Email task queued successfully")
	})

	log.Println("Email task producer is running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
