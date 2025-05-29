package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// Подписка на тему "updates"
	sub, err := nc.Subscribe("updates", func(msg *nats.Msg) {
		log.Printf("Received message on %s: %s", msg.Subject, string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}
	defer sub.Unsubscribe()

	// Явное ожидание подтверждения подписки
	if err := nc.Flush(); err != nil {
		log.Fatalf("Error flushing NATS connection: %v", err)
	}

	if err := sub.AutoUnsubscribe(1000); err != nil {
		log.Printf("Warning: Failed to set auto-unsubscribe: %v", err)
	}

	log.Printf("Subscribed to 'updates'. Listening for messages...")

	// Настройка сигналов для правильного завершения
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала или периодическая проверка состояния
	for {
		select {
		case sig := <-sigCh:
			log.Printf("Received signal %v, shutting down...", sig)
			return
		case <-time.After(10 * time.Second):
			// Периодическая проверка состояния соединения
			if nc.Status() != nats.CONNECTED {
				log.Printf("Warning: NATS connection status: %v", nc.Status())
			} else {
				log.Printf("NATS connection still active. Waiting for messages...")
			}
		}
	}
}
