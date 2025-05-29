# Redis Pub/Sub

Пример работы pub и sub

## Запуск

1. Запуск Redis:
`docker run --name redis -p 6379:6379 -d redis`

2. В 1 терминале запускаем подписчика:
`cd subscriber`
`go run main.go`

3. В другом терминале запускаем публикатора:
`cd publisher`
`go run main.go`

## Как работает

Паб отправляет сообщение в канал каждые 2 секунды
Саб прослушивает канал и получает сообщения

Паб
![alt text](screenshots/pub.png)

Саб
![alt text](screenshots/sub.png)
