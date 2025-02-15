

build:
	docker build .
	docker compose up -d 

restart:
	docker compose down -v
	docker compose up --build -d

clean:
	docker compose down -v

test:
	go test  -v ./internal/shortener

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/proto/shortener/link_shortening.proto

migrations:
	export GOOSE_DBSTRING="user=postgres dbname=postgres sslmode=disable password=password"