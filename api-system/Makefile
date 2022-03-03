.PHONY: all swag clean upgrade check build-image run help run-dev-env

all: swag run

swag:
	@echo "latest swag tool required"
	swag init -o ./api/docs -g ./cmd/app/main.go
	swag fmt -g ./cmd/app/main.go

clean:
	go clean

upgrade:
	go get -u ./cmd/app/

check:
	go fmt ./cmd/app/
	go vet ./cmd/app/

build-image:
	@echo "docker required"
	docker build . -t axiangcoding/antonstar/api-system:latest

run:
	go run ./cmd/app/main.go



