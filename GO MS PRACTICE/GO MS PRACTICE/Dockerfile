FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy # Убедитесь, что все зависимости установлены

COPY . .

RUN go build -o my-go-app

RUN ls -la /app/my-go-app # Проверка: существует ли файл?



FROM alpine:latest

RUN apk add --no-cache libc6-compat # Для совместимости, если нужно

WORKDIR /usr/local/bin/

COPY --from=builder /app/my-go-app .

CMD ["my-go-app"]