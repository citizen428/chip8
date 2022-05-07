PRG := chip8

run: build
	@./bin/${PRG} -scaleFactor=15

build:
	@go build -o bin/${PRG} ./cmd/${PRG}

test:
	@go test ./internal/emulator

lint:
	@go vet ./...
	@staticcheck ./...
	@shadow ./...

clean:
	@go clean
	@rm bin/${PRG}

.PHONY: build run lint clean
