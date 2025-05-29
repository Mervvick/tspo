package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Подключение к NATS серверу с обработчиками событий
	nc, err := nats.Connect("nats://localhost:4222",
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			log.Printf("NATS Error: %v", err)
		}),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS Disconnected: %v", nc.LastError())
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS Reconnected: %v", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("NATS Connection Closed: %v", nc.LastError())
		}),
	)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Close()

	if nc.Status() != nats.CONNECTED {
		log.Fatalf("Failed to connect to NATS. Status: %v", nc.Status())
	}
	log.Printf("Connected to NATS server at %s", nc.ConnectedUrl())

	// Публикация сообщений
	counter := 1
	subject := "updates"

	for {
		message := fmt.Sprintf("Update #%d at %s", counter, time.Now().Format(time.RFC3339))

		err = nc.Publish(subject, []byte(message))
		if err != nil {
			log.Printf("Error publishing message: %v", err)
		} else {
			log.Printf("Published message to %s: %s", subject, message)
		}

		// Явное принудительное отправление всех буферизованных сообщений
		if err := nc.Flush(); err != nil {
			log.Printf("Error flushing NATS connection: %v", err)
		}

		counter++
		time.Sleep(3 * time.Second)
	}
}
