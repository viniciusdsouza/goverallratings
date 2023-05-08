.PHONY: default run build test docs clean

APP_NAME=goverallratings

default: run-with-docs

run-with-docs:
	@echo "STARTING APPLICATION WITH DOCS..."
	@swag init
	@go run main.go
run:
	@echo "STARTING APPLICATION..."
	@go run main.go
build:
	@echo "STARTING BUILD..."
	@go build -o $(APP_NAME) main.go
test:
	@echo "STARTING TESTS..."
	@go test ./ ...
docs:
	@echo "STARTING SWAG GENERATION..."
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs