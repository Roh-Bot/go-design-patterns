.PHONY: all build run deps

all:
	@make run

build:
	@go build --race main.go

run:
	@echo "Running the application"
	@go run --race main.go

deps:
	@echo "Fetching dependencies"
	@go mod tidy
	@echo "Updating dependencies"
	@go get -u
