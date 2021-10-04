.PHONY: all install build run serve-swagger

all: run

install:
	go mod tidy

build:
	go build -o bin/main ./cmd/main.go

run:
	go run ./cmd/main.go

check-swagger:
	which swagger || (go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger:
	swagger generate spec -o ./docs/swagger.yaml --scan-models

serve-swagger: swagger
	swagger serve -F=swagger ./docs/swagger.yaml -F redoc