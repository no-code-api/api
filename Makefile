build:
	@go build -o bin/no-code-api

run: build
	@./bin/no-code-api
