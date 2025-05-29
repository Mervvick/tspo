package models

import (
	"encoding/json"
	"time"
)

type OrderStatus string

const (
	StatusPending  OrderStatus = "PENDING"
	StatusPaid     OrderStatus = "PAID"
	StatusCanceled OrderStatus = "CANCELED"
)

// --- Events ---
type EventType string

const (
	EventOrderCreated  EventType = "OrderCreated"
	EventOrderPaid     EventType = "OrderPaid"
	EventOrderCanceled EventType = "OrderCanceled"
)

type Event struct {
	Type      EventType       `json:"type"`
	OrderID   string          `json:"order_id"`
	Timestamp time.Time       `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

// --- Read model ---
type Order struct {
	ID     string      `json:"id"`
	Status OrderStatus `json:"status"`
}
