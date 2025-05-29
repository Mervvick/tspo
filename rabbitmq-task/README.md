# RabbitMQ Task Queue Example

This example demonstrates using RabbitMQ as a message broker for handling asynchronous email sending tasks.

## Setup

1. Start RabbitMQ:
docker run -d --hostname my-rabbit --name rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management

2. Run the consumer in one terminal:
cd consumer
go run main.go

3. Run the producer in another terminal:
cd producer
go run main.go

## Testing

Send a test email task:
curl -X POST http://localhost:8081/send-email
-H "Content-Type: application/json"
-d '{
"to": "user@example.com",
"subject": "Test Email",
"body": "This is a test email from RabbitMQ task queue"
}'

## How it works

1. The producer exposes an HTTP endpoint that accepts email task details
2. When a request is received, it serializes the task and pushes it to a RabbitMQ queue
3. The consumer listens to the queue and processes email tasks as they arrive
4. The consumer acknowledges messages only after successful processing