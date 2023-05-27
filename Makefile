BUILD_DIR=.build
.PHONY: default run build test docs clean build-deps build-webserver

APP_NAME=goverallratings

default: run-with-docs

run-with-docs:
	@echo "STARTING APPLICATION WITH DOCS..."
	@swag init
	@go run main.go

run:
	@echo "STARTING APPLICATION..."
	@go run main.go

test:
	@echo "STARTING TESTS..."
	@go test ./ ...

mk-build-dir:
	@mkdir -p ${BUILD_DIR}

docs:
	@echo "STARTING SWAG GENERATION..."
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs

build-deps:
	@go get -d -v ./...
	@go mod tidy

build: build-deps build-webserver

build-webserver:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o ./${BUILD_DIR}/webserver .
