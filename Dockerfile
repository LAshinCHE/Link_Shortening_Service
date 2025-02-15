FROM golang:latest


RUN go install github.com/pressly/goose/v3/cmd/goose@latest


RUN mkdir /app
WORKDIR /app


COPY . .

EXPOSE 8080


RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/app/main.go

ENTRYPOINT ["sh", "-c", "./app -dbtype $STORAGE_TYPE"]