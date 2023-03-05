APP=fizzbuzz
VERSION=0.1.0
BIN_FOLDER=bin
BIN_NAME=${BIN_FOLDER}/fizzbuzz

install:
	@echo "Installing dependencies"
	@go get .

build: clean install
	@echo "Compiling binary"
	@go build -o ${BIN_NAME} .

start:
	@echo "Starting ${BIN_NAME}"
	@echo
	@${BIN_NAME}

dev:
	@echo "Running dev mode. Log level to 'DEBUG'"
	@DEBUG_LEVEL=DEBUG go run .

test:
	@echo "Running tests"
	@go test -v .

clean:
	@echo "Cleaning project"
	@go clean
	@rm -rf ${BIN_FOLDER}

docker:
	@echo "Building Dockerfile"
	@docker build -t edjubert/leboncoin .

all: build start

.PHONY: build clean run dev install
