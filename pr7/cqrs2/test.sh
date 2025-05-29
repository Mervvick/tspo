#!/bin/bash

echo "Testing Order API..."

# Создание заказа
ORDER_ID=$(curl -s -X POST http://localhost:8080/orders | jq -r '.order_id')
echo "Created order with ID: $ORDER_ID"

# Проверка статуса заказа
STATUS=$(curl -s http://localhost:8080/orders/$ORDER_ID | jq -r '.status')
echo "Order status: $STATUS"
if [ "$STATUS" != "PENDING" ]; then
  echo "ERROR: Expected PENDING status"
  exit 1
fi

# Оплата заказа
curl -s -X POST http://localhost:8080/orders/$ORDER_ID/pay
echo "Paid for order"

# Проверка статуса после оплаты
STATUS=$(curl -s http://localhost:8080/orders/$ORDER_ID | jq -r '.status')
echo "Order status after payment: $STATUS"
if [ "$STATUS" != "PAID" ]; then
  echo "ERROR: Expected PAID status"
  exit 1
fi

# Просмотр всех событий
echo "Events:"
curl -s http://localhost:8080/events | jq '.'

echo "All tests passed!"
