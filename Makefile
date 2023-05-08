.PHONY: default run build test docs clean

APP_NAME=goverallratings

default: run

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