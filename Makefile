all: dependencies test build

ci: dependencies test-ci build

build:
	mkdir -p build
	go build -o ./build/glitzz ./cmd/glitzz

run:
	./build/glitzz

doc:
	@echo "http://localhost:6060/pkg/github.com/lovelaced/glitzz/"
	godoc -http=:6060

test:
	go test ./...

test-ci:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

test-verbose:
	go test -v ./...

test-short:
	go test -short ./...

clean:
	rm -f ./build/glitzz

dependencies:
	go get -t ./...

docker:
	@echo 
	@echo "    If you are creating the image for the first time and want to create a default config run:"
	@echo "        $$ make docker-init"
	@echo "        $$ cat _data/config.json"
	@echo 
	@echo 
	@echo "    If you want to update the image without creating the config run:"
	@echo "        $$ git pull"
	@echo "        $$ make docker-build"
	@echo 
	@echo 
	@echo "    If you want to stop the image run:"
	@echo "        $$ make docker-stop"
	@echo 
	@echo 
	@echo "    If you want to start the image run:"
	@echo "        $$ make docker-start"
	@echo 
	@echo 

docker-init: docker-create-directories docker-build docker-create-config

docker-create-directories:
	mkdir -p _data

docker-create-config:
	sudo docker-compose run --rm glitzz glitzz default_config >> _data/config.json

docker-build:
	sudo docker-compose build
	sudo docker-compose up --no-start

docker-start:
	sudo docker-compose start

docker-stop:
	sudo docker-compose stop

.PHONY: all ci build run doc test test-ci test-verbose test-short clean dependencies docker docker-init docker-create-directories docker-create-config docker-build docker-start docker-stop
