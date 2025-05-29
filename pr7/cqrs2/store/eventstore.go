package store

import (
	"sync"

	"order-service/models"
)

type EventStore struct {
	eventLog []models.Event
	orders   map[string]models.Order
	mutex    sync.RWMutex
}

func NewEventStore() *EventStore {
	return &EventStore{
		eventLog: []models.Event{},
		orders:   make(map[string]models.Order),
	}
}

func (es *EventStore) AppendEvent(e models.Event) {
	es.mutex.Lock()
	defer es.mutex.Unlock()

	es.eventLog = append(es.eventLog, e)
	es.applyEvent(e)
}

func (es *EventStore) applyEvent(e models.Event) {
	switch e.Type {
	case models.EventOrderCreated:
		es.orders[e.OrderID] = models.Order{ID: e.OrderID, Status: models.StatusPending}
	case models.EventOrderPaid:
		if o, ok := es.orders[e.OrderID]; ok {
			o.Status = models.StatusPaid
			es.orders[e.OrderID] = o
		}
	case models.EventOrderCanceled:
		if o, ok := es.orders[e.OrderID]; ok {
			o.Status = models.StatusCanceled
			es.orders[e.OrderID] = o
		}
	}
}

func (es *EventStore) GetOrder(orderID string) (models.Order, bool) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()

	order, exists := es.orders[orderID]
	return order, exists
}

func (es *EventStore) GetAllEvents() []models.Event {
	es.mutex.RLock()
	defer es.mutex.RUnlock()

	// Создаем копию списка событий
	events := make([]models.Event, len(es.eventLog))
	copy(events, es.eventLog)
	return events
}

func (es *EventStore) RebuildState() {
	es.mutex.Lock()
	defer es.mutex.Unlock()

	// Очищаем текущее состояние заказов
	es.orders = make(map[string]models.Order)

	// Применяем все события заново
	for _, e := range es.eventLog {
		es.applyEvent(e)
	}
}
