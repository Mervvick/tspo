package main

// EventStore хранит все события и управляет ими
type EventStore struct {
	Events []Event
}

// NewEventStore создает новый экземпляр EventStore
func NewEventStore() *EventStore {
	return &EventStore{
		Events: make([]Event, 0),
	}
}

// AddEvents добавляет события в хранилище
func (es *EventStore) AddEvents(events []Event) {
	es.Events = append(es.Events, events...)
}

// GetEventsByOrderID возвращает все события для определенного заказа
func (es *EventStore) GetEventsByOrderID(orderID string) []Event {
	result := make([]Event, 0)
	for _, event := range es.Events {
		if event.GetOrderID() == orderID {
			result = append(result, event)
		}
	}
	return result
}

// GetOrderByID восстанавливает заказ из событий и возвращает его
func (es *EventStore) GetOrderByID(orderID string) *Order {
	events := es.GetEventsByOrderID(orderID)
	if len(events) == 0 {
		return nil
	}

	order := &Order{ID: orderID}
	for _, event := range events {
		applyEvent(order, event)
	}
	return order
}

// GetAllOrders восстанавливает и возвращает все заказы из событий
func (es *EventStore) GetAllOrders() []*Order {
	// Создаем map для быстрого поиска заказов по ID
	orderMap := make(map[string]*Order)

	// Проходимся по всем событиям и применяем их к соответствующим заказам
	for _, event := range es.Events {
		orderID := event.GetOrderID()

		order, exists := orderMap[orderID]
		if !exists {
			order = &Order{ID: orderID}
			orderMap[orderID] = order
		}
		applyEvent(order, event)
	}

	// Преобразуем map в slice для возврата
	orders := make([]*Order, 0, len(orderMap))
	for _, order := range orderMap {
		orders = append(orders, order)
	}

	return orders
}
