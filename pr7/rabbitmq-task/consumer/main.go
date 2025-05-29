package main

import (
	"encoding/json"
	"log"
	"time"

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

// Имитация отправки email
func sendEmail(task EmailTask) error {
	log.Printf("Sending email to: %s", task.To)
	log.Printf("Subject: %s", task.Subject)
	log.Printf("Body: %s", task.Body)

	// Имитация задержки при отправке email
	time.Sleep(1 * time.Second)

	log.Printf("Email sent successfully to %s", task.To)
	return nil
}

func main() {
	// Подключение к RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Объявление очереди (для гарантии, что она существует)
	q, err := ch.QueueDeclare(
		"email_tasks", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Установка prefetch count, чтобы получать сообщения по одному
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// Получение сообщений
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var task EmailTask
			if err := json.Unmarshal(d.Body, &task); err != nil {
				log.Printf("Error parsing task: %s", err)
				d.Nack(false, false) // отклонить сообщение
				continue
			}

			// Обработка задачи
			err := sendEmail(task)
			if err != nil {
				log.Printf("Error processing task: %s", err)
				d.Nack(false, true) // вернуть сообщение в очередь
				continue
			}

			d.Ack(false) // подтвердить успешную обработку
			log.Printf("Task processed successfully")
		}
	}()

	log.Printf(" [*] Email consumer waiting for messages. To exit press CTRL+C")
	<-forever
}
