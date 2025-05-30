FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod init consumer && \
    go get github.com/nats-io/nats.go && \
    go build -o consumer main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/consumer .
CMD ["./consumer"]
