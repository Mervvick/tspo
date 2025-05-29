package main

import (
	"time"
)

// Event представляет собой интерфейс для всех событий в системе
type Event interface {
	GetID() string
	GetTimestamp() time.Time
	GetType() string
	GetOrderID() string // Добавим метод для получения OrderID
}

// BaseEvent содержит общие поля для всех событий
type BaseEvent struct {
	ID        string    // ID события
	OrderID   string    // ID заказа, к которому относится событие
	Timestamp time.Time // Время создания события
}

func (e BaseEvent) GetID() string {
	return e.ID
}

func (e BaseEvent) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e BaseEvent) GetOrderID() string {
	return e.OrderID
}

// BaseEvent не может реализовать GetType(), так как это специфично для каждого типа события

// OrderCreatedEvent - событие создания заказа
type OrderCreatedEvent struct {
	BaseEvent
	CustomerID string
	Items      []OrderItem
}

func (e OrderCreatedEvent) GetType() string {
	return "OrderCreated"
}

// OrderPaidEvent - событие оплаты заказа
type OrderPaidEvent struct {
	BaseEvent
	PaymentID string
	Amount    float64
}

func (e OrderPaidEvent) GetType() string {
	return "OrderPaid"
}

// OrderCancelledEvent - событие отмены заказа
type OrderCancelledEvent struct {
	BaseEvent
	Reason string
}

func (e OrderCancelledEvent) GetType() string {
	return "OrderCancelled"
}
