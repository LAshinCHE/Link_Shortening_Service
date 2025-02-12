include .env
ENV_FILE := .env


# build docker image
build:
	docker-compose build

# migrations
migrate:
	docker-compose --env-file $(ENV_FILE) run --rm app sh -c 'goose -dir db/migrations postgres "$(DATABASE_URL)" up'

# app and postgrace
up-all:
	docker-compose up -d postgres app

down:
	docker-compose down


# db
up-db:
	docker-compose up -d postgres

stop-db:
	docker-compose stop postgres

start-db:
	docker-compose start postgres

down-db:
	docker-compose down postgres

#app
up-app:
	docker-compose up -d app

stop-app:
	docker-compose stop app

start-app:
	docker-compose start app

down-app:
	docker-compose down app