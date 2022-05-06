PRG := chip8

run: build
	@./bin/${PRG}

build:
	@go build -o bin/${PRG} ./cmd/${PRG}

lint:
	@go vet ./...
	@staticcheck ./...
	@shadow ./...

clean:
	@go clean
	@rm bin/${PRG}

.PHONY: build run lint clean
