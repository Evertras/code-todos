GO_FILES=$(shell find . -name '*.go' -not -path './vendor/*')

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

todos.md: $(GO_FILES)
	go run cmd/code-todos/main.go find internal cmd > todos.md

bin/code-todos: $(GO_FILES)
	@mkdir -p bin
	go build -o bin/code-todos cmd/code-todos/main.go

./git/hooks/pre-commit: .evertras/pre-commit.sh
	cp .evertras/pre-commit.sh .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
