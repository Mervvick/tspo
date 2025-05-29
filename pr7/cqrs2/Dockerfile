# Используем многоступенчатую сборку
FROM golang:1.20-alpine AS builder

# Установка необходимых пакетов
RUN apk add --no-cache git

# Создание директории для приложения
WORKDIR /app

# Копирование файлов зависимостей
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o order-service .

# Финальный образ
FROM alpine:latest

# Установка зависимостей и настройка для работы с https
RUN apk --no-cache add ca-certificates tzdata

# Копирование скомпилированного приложения
COPY --from=builder /app/order-service /order-service

# Настройка переменной окружения
ENV PORT=8080

# Экспонирование порта
EXPOSE 8080

# Запуск приложения
CMD ["/order-service"]
