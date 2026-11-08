BINARY := chronicle
MAIN := ./cmd/chronicle
BIN_DIR := ./bin

.PHONY: build install test lint clean

build:
	mkdir -p $(BIN_DIR)
	go build -trimpath -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY) $(MAIN)

install:
	go install $(MAIN)

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -rf $(BIN_DIR)
