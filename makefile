.PHONY: dep run build migrate test test-ci build-docker serve-docker

run:
	go run main.go

migrate-up:
	 migrate -database "postgres://fajarfathirwan@localhost:5432/pos?sslmode=disable" -source "file:///Users/fajarfathirwan/Project/go-docker/migrations" up

migrate-down:
	 migrate -database "postgres://fajarfathirwan@localhost:5432/pos?sslmode=disable" -source "file:///Users/fajarfathirwan/Project/go-docker/migrations" up

dep:
	go mod download
	go mod verify

build:
	go build .

lint:
	go fmt ./...
	golangci-lint run ./...

test:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

test-ci:
	go test -cover -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

build-docker:
	docker build -t go-docker .

serve-docker:
	docker run -d -p 8081:8081 go-docker
