NAME=gocloud

SRC_DIR=./cmd
TEMP_DIR=./tmp
BUILD_DIR=./build
DOCS_DIR=./docs
SERVER_FILE=./internal/application/server/server.go

dev:
	air -c .air.toml

run: build
	$(BUILD_DIR)/$(NAME)

build: clean swagger
	go build -o $(BUILD_DIR)/$(NAME) $(SRC_DIR)/main.go
	cp .env $(BUILD_DIR)

test:
	go test ./...

fmt:
	go fmt ./...

deps:
	go mod tidy

clean:
	rm -rf $(TEMP_DIR) $(DOCS_DIR) $(BUILD_DIR)

swagger:
	swag init -g $(SERVER_FILE)