todos.md: $(shell find . -name '*.go')
	go run cmd/code-todos/main.go find internal cmd > todos.md

.PHONY: try
try:
	go run cmd/code-todos/main.go find internal cmd

.PHONY: test
test:
	go test ./internal/todos
