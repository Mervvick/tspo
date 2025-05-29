package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"order-service/models"
	"order-service/store"
)

type Handlers struct {
	eventStore *store.EventStore
}

func NewHandlers(es *store.EventStore) *Handlers {
	return &Handlers{
		eventStore: es,
	}
}

// --- Command Handlers ---
func (h *Handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	orderID := uuid.New().String()
	event := models.Event{
		Type:      models.EventOrderCreated,
		OrderID:   orderID,
		Timestamp: time.Now(),
		Data:      json.RawMessage(`{}`),
	}

	h.eventStore.AppendEvent(event)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"order_id": orderID})
}

func (h *Handlers) PayOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]

	// Проверяем существование заказа
	if _, exists := h.eventStore.GetOrder(orderID); !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	event := models.Event{
		Type:      models.EventOrderPaid,
		OrderID:   orderID,
		Timestamp: time.Now(),
		Data:      json.RawMessage(`{}`),
	}

	h.eventStore.AppendEvent(event)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) CancelOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]

	// Проверяем существование заказа
	if _, exists := h.eventStore.GetOrder(orderID); !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	event := models.Event{
		Type:      models.EventOrderCanceled,
		OrderID:   orderID,
		Timestamp: time.Now(),
		Data:      json.RawMessage(`{}`),
	}

	h.eventStore.AppendEvent(event)
	w.WriteHeader(http.StatusNoContent)
}

// --- Query Handlers ---
func (h *Handlers) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]
	order, exists := h.eventStore.GetOrder(orderID)

	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *Handlers) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events := h.eventStore.GetAllEvents()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// --- Middleware ---
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логируем запрос
		println(time.Now().Format(time.RFC3339), r.Method, r.RequestURI)

		// Передаем запрос дальше
		next.ServeHTTP(w, r)
	})
}
