.PHONY: default
default: ./git/hooks/pre-commit todos.md

.PHONY: build
build: bin/code-todos

.PHONY: try
try:
	go run cmd/code-todos/main.go find internal cmd

.PHONY: test
test:
	go test ./internal/todos

todos.md: $(shell find . -name '*.go')
	go run cmd/code-todos/main.go find internal cmd > todos.md

bin/code-todos: $(shell find . -name '*.go')
	@mkdir -p bin
	go build -o bin/code-todos cmd/code-todos/main.go

./git/hooks/pre-commit: .evertras/pre-commit.sh
	cp .evertras/pre-commit.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
