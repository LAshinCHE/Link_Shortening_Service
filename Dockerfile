# Используем официальный образ Go для сборки
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/link-shortener ./cmd/main.go

# Используем минимальный образ Alpine для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарный файл из стадии сборки
COPY --from=builder /app/bin/link-shortener ./link-shortener

# Копируем миграции
COPY /migrations ./db/migrations

# Устанавливаем Goose для выполнения миграций
RUN apk add --no-cache postgresql-client && \
    wget -O /usr/local/bin/goose https://github.com/pressly/goose/releases/download/v3.15.1/goose_linux_x86_64 && \
    chmod +x /usr/local/bin/goose

# Указываем порт, который будет использовать приложение
EXPOSE 8080

# Команда для запуска приложения
CMD ["./link-shortener"]