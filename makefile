.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

// DEV
migrate -path migrations -database "postgres://localhost/restapi_dev?sslmode=disable" up

// TEST
migrate -path migrations -database "postgres://localhost/restapi_test?sslmode=disable" up