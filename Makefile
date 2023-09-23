.PHONY: default
default:
	go run cmd/code-todos/main.go find internal cmd

.PHONY: test
test:
	go test ./internal/todos
