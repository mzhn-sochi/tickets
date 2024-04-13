BIN_DIR=bin

MODULE=tickets

build:
	mkdir -p $(BIN_DIR) #
	go mod tidy
	go build -o $(BIN_DIR)/ -v ./cmd/service

gen:
	protoc --go_opt=module=$(MODULE) --go-grpc_opt=module=$(MODULE) \
	 	--go_out=. --go-grpc_out=. \
		-I ./proto ./proto/share.proto ./proto/tickets.proto
	wire ./internal/app

deploy:
	make gen
	docker compose build
	docker compose down
	docker compose up -d

migrate.up:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5434/tickets?sslmode=disable' up

migrate.down:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5434/tickets?sslmode=disable' down

clean:
	rm -r $(BIN_DIR) api #
