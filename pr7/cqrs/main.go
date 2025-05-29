package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем хранилище событий и обработчик команд
	eventStore := NewEventStore()
	commandHandler := NewCommandHandler(eventStore)

	// Создаем заказ
	orderID := generateID()
	createCmd := CreateOrderCommand{
		OrderID:    orderID,
		CustomerID: "customer-123",
		Items: []OrderItem{
			{ProductID: "product-1", Quantity: 2, Price: 10.0},
			{ProductID: "product-2", Quantity: 1, Price: 20.0},
		},
	}
	err := commandHandler.Handle(createCmd)
	if err != nil {
		fmt.Printf("Ошибка при создании заказа: %v\n", err)
	} else {
		fmt.Printf("Заказ %s создан\n", orderID)
	}

	// Оплачиваем заказ
	payCmd := PayOrderCommand{
		OrderID:   orderID,
		PaymentID: "payment-456",
		Amount:    40.0,
	}
	err = commandHandler.Handle(payCmd)
	if err != nil {
		fmt.Printf("Ошибка при оплате заказа: %v\n", err)
	} else {
		fmt.Printf("Заказ %s оплачен\n", orderID)
	}

	// Пытаемся отменить оплаченный заказ (должна быть ошибка)
	cancelCmd := CancelOrderCommand{
		OrderID: orderID,
		Reason:  "Передумал",
	}
	err = commandHandler.Handle(cancelCmd)
	if err != nil {
		fmt.Printf("Ошибка при отмене заказа: %v\n", err)
	} else {
		fmt.Printf("Заказ %s отменен\n", orderID)
	}

	// Создаем еще один заказ
	anotherOrderID := generateID()
	createCmd2 := CreateOrderCommand{
		OrderID:    anotherOrderID,
		CustomerID: "customer-456",
		Items: []OrderItem{
			{ProductID: "product-3", Quantity: 1, Price: 30.0},
		},
	}
	err = commandHandler.Handle(createCmd2)
	if err != nil {
		fmt.Printf("Ошибка при создании заказа: %v\n", err)
	} else {
		fmt.Printf("Заказ %s создан\n", anotherOrderID)
	}

	// Отменяем второй заказ
	cancelCmd2 := CancelOrderCommand{
		OrderID: anotherOrderID,
		Reason:  "Товар закончился",
	}
	err = commandHandler.Handle(cancelCmd2)
	if err != nil {
		fmt.Printf("Ошибка при отмене заказа: %v\n", err)
	} else {
		fmt.Printf("Заказ %s отменен\n", anotherOrderID)
	}

	// Получаем и выводим все заказы
	orders := eventStore.GetAllOrders()
	fmt.Println("\nСписок всех заказов:")
	for _, order := range orders {
		fmt.Printf("Заказ %s (Клиент: %s, Статус: %s)\n",
			order.ID, order.CustomerID, order.Status)
	}

	// Выводим все события
	fmt.Println("\nЖурнал событий:")
	for i, event := range eventStore.Events {
		fmt.Printf("%d. %s: %s (Заказ: %s, Время: %s)\n",
			i+1, event.GetType(), event.GetID(),
			event.GetOrderID(), event.GetTimestamp().Format(time.RFC3339))
	}
}

// generateID генерирует уникальный ID
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
