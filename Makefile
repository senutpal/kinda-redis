BINARY := kinda-redis
BIN_DIR := bin
BIN := $(BIN_DIR)/$(BINARY)

GO := go
GOCMD := $(GO)

.PHONY: build run test fmt clean lint tidy

build: $(BIN_DIR)
	@$(GOCMD) build -o $(BIN) src/main.go

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

run: build
	@$(BIN)

test:
	@$(GO) test -v ./...

fmt:
	@$(GO) fmt ./...

lint:
	@golangci-lint run ./...

tidy:
	@$(GO) mod tidy

clean:
	@rm -rf $(BIN_DIR)