package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"order-service/handlers"
	"order-service/store"
)

func main() {
	// Инициализация хранилища
	eventStore := store.NewEventStore()

	// Восстановление состояния
	eventStore.RebuildState()

	// Инициализация обработчиков
	h := handlers.NewHandlers(eventStore)

	// Настройка маршрутизации
	r := mux.NewRouter()

	// Добавляем middleware для логирования
	r.Use(handlers.LoggingMiddleware)

	// Команды
	r.HandleFunc("/orders", h.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}/pay", h.PayOrder).Methods("POST")
	r.HandleFunc("/orders/{id}/cancel", h.CancelOrder).Methods("POST")

	// Запросы
	r.HandleFunc("/orders/{id}", h.GetOrder).Methods("GET")
	r.HandleFunc("/events", h.GetAllEvents).Methods("GET")

	// Обработка сигналов остановки
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине
	go func() {
		port := getEnv("PORT", "8080")
		log.Printf("Server starting on :%s\n", port)
		if err := http.ListenAndServe(":"+port, r); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Ожидание сигнала остановки
	<-stop
	log.Println("Server shutting down...")
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
