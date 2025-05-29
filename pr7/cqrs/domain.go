package main

import (
	"time"
)

// OrderItem представляет товар в заказе
type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

// Order представляет заказ в системе
type Order struct {
	ID         string
	CustomerID string
	Items      []OrderItem
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PaymentID  string
	Amount     float64
}

// applyEvent применяет событие к заказу, обновляя его состояние
func applyEvent(order *Order, event Event) {
	switch e := event.(type) {
	case OrderCreatedEvent:
		order.CustomerID = e.CustomerID
		order.Items = e.Items
		order.Status = "Created"
		order.CreatedAt = e.Timestamp
		order.UpdatedAt = e.Timestamp
	case OrderPaidEvent:
		order.Status = "Paid"
		order.UpdatedAt = e.Timestamp
		order.PaymentID = e.PaymentID
		order.Amount = e.Amount
	case OrderCancelledEvent:
		order.Status = "Cancelled"
		order.UpdatedAt = e.Timestamp
	}
}
