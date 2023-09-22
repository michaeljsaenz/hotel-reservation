BINARY_NAME=api

build:
	rm -rf bin/${BINARY_NAME}
	@go build -o bin/api

run: build
	@./bin/api --listenAddr  :5000

test:
	@go test -v ./...