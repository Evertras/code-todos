.PHONY: default
default:
	go run cmd/code-todos/main.go find .

.PHONY: test
test:
	go test ./internal/todos
