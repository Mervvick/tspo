package main

import (
	"fmt"
	"time"
)

// Command представляет собой интерфейс для всех команд в системе
type Command interface {
	Execute(es *EventStore) ([]Event, error)
}

// CreateOrderCommand - команда для создания заказа
type CreateOrderCommand struct {
	OrderID    string
	CustomerID string
	Items      []OrderItem
}

// Execute выполняет команду создания заказа
func (c CreateOrderCommand) Execute(es *EventStore) ([]Event, error) {
	event := OrderCreatedEvent{
		BaseEvent: BaseEvent{
			ID:        generateID(),
			OrderID:   c.OrderID,
			Timestamp: time.Now(),
		},
		CustomerID: c.CustomerID,
		Items:      c.Items,
	}
	return []Event{event}, nil
}

// PayOrderCommand - команда для оплаты заказа
type PayOrderCommand struct {
	OrderID   string
	PaymentID string
	Amount    float64
}

// Execute выполняет команду оплаты заказа
func (c PayOrderCommand) Execute(es *EventStore) ([]Event, error) {
	// Проверяем, существует ли заказ и можно ли его оплатить
	order := es.GetOrderByID(c.OrderID)
	if order == nil {
		return nil, fmt.Errorf("заказ с ID %s не найден", c.OrderID)
	}
	if order.Status != "Created" {
		return nil, fmt.Errorf("заказ с ID %s нельзя оплатить (текущий статус: %s)", c.OrderID, order.Status)
	}

	event := OrderPaidEvent{
		BaseEvent: BaseEvent{
			ID:        generateID(),
			OrderID:   c.OrderID,
			Timestamp: time.Now(),
		},
		PaymentID: c.PaymentID,
		Amount:    c.Amount,
	}
	return []Event{event}, nil
}

// CancelOrderCommand - команда для отмены заказа
type CancelOrderCommand struct {
	OrderID string
	Reason  string
}

// Execute выполняет команду отмены заказа
func (c CancelOrderCommand) Execute(es *EventStore) ([]Event, error) {
	// Проверяем, существует ли заказ и можно ли его отменить
	order := es.GetOrderByID(c.OrderID)
	if order == nil {
		return nil, fmt.Errorf("заказ с ID %s не найден", c.OrderID)
	}
	if order.Status == "Cancelled" {
		return nil, fmt.Errorf("заказ с ID %s уже отменен", c.OrderID)
	}
	if order.Status == "Paid" {
		return nil, fmt.Errorf("оплаченный заказ нельзя отменить")
	}

	event := OrderCancelledEvent{
		BaseEvent: BaseEvent{
			ID:        generateID(),
			OrderID:   c.OrderID,
			Timestamp: time.Now(),
		},
		Reason: c.Reason,
	}
	return []Event{event}, nil
}
