FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod init publisher && \
    go get github.com/nats-io/nats.go && \
    go build -o publisher main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/publisher .
CMD ["./publisher"]
