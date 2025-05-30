package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
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

// Имитация отправки email с возможными ошибками
func sendEmail(task EmailTask) error {
	log.Printf("Attempting to send email to: %s", task.To)
	log.Printf("Subject: %s", task.Subject)
	log.Printf("Body: %s", task.Body)

	// Имитация задержки при отправке email
	time.Sleep(1 * time.Second)

	// Имитация случайных ошибок (примерно в 20% случаев)
	if rand.Float32() < 0.2 {
		log.Printf("ERROR: Failed to send email to %s (simulated failure)", task.To)
		return errors.New("simulated email sending failure")
	}

	log.Printf("Email sent successfully to %s", task.To)
	return nil
}

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

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
		1,     // prefetch count - обрабатываем только одно сообщение за раз
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// Получение сообщений с отключенным автоподтверждением (auto-ack: false)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack = false - ВАЖНО: ручное подтверждение!
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Println("Email consumer started. Waiting for messages...")
	log.Println("Press CTRL+C to exit")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var task EmailTask
			if err := json.Unmarshal(d.Body, &task); err != nil {
				log.Printf("Error parsing task: %s", err)

				// Отклоняем сообщение с requeue=false,
				// так как оно некорректно и повторная обработка не поможет
				d.Reject(false)
				continue
			}

			// Обработка задачи
			err := sendEmail(task)
			if err != nil {
				log.Printf("Error processing task: %s", err)

				// Отклоняем сообщение с requeue=true,
				// чтобы оно вернулось в очередь для повторной обработки
				d.Reject(true)

				// Добавляем задержку перед обработкой следующего сообщения,
				// чтобы избежать быстрого цикла повторных обработок при ошибках
				time.Sleep(3 * time.Second)
				continue
			}

			// Подтверждаем успешную обработку
			d.Ack(false) // multiple=false - подтверждаем только текущее сообщение
			log.Printf("Task processed successfully and acknowledged")
		}
	}()

	<-forever
}
