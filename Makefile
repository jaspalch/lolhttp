BIN=lolhttp

.PHONY: run test clean
build: test
	@go build -o $(BIN) ./cmd/$(BIN)/

run: build
	@./$(BIN)

test:
	@go test ./...

clean:
	@rm -f $(BIN)
