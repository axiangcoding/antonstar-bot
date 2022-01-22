.PHONY: all swag clean upgrade check build-image run help run-dev-env

all: swag run

swag:
	@echo "latest swag tool required"
	swag init --output ./api/docs
	swag fmt

clean:
	go clean

upgrade:
	go get -u

check:
	go fmt ./
	go vet ./

build-image:
	@echo "docker required"
	docker build . -t axiangcoding/antonstar/api-system:latest

run:
	go run ./



