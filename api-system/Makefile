.PHONY: gen-swag clean upgrade check build-image run-app help

gen-swag:
	@echo "generate swagger files..."
	swag init -o ./api -g internal/controller/http/v1/router.go
	swag fmt -g internal/controller/http/v1/router.go

clean:
	@echo "go clean"
	go clean

upgrade:
	@echo "go upgrade go.mod"
	go get -u ./cmd/app
	go get -u ./cmd/dal

check:
	go fmt ./cmd/app
	go vet ./cmd/app

build-image:
	@echo "docker required"
	docker build --build-arg=VERSION=latest . -t axiangcoding/antonstar-bot-server:latest

run-app:
	go run ./cmd/app

help:
	@echo "Please take a look at Makefile"