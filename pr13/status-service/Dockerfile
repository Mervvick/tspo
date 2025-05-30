FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod init status-service && \
    go get github.com/nats-io/nats.go && \
    go build -o status-service main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/status-service .
CMD ["./status-service"]
