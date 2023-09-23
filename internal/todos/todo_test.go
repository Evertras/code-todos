package todos_test

import (
	"testing"

	"github.com/evertras/code-todos/internal/todos"
	"github.com/stretchr/testify/assert"
)

func TestFindTodosInFile(t *testing.T) {
	found, errs := todos.FindTodos("testdata/go/main.go")

	assert.Len(t, errs, 0)
	assert.Len(t, found, 1)

	expectedTodo := todos.Todo{
		Filename:    "testdata/go/main.go",
		PackageName: "main",
		Line:        4,
		Text:        "TODO: Write a really cool thing here.\nAnd do it on multiple lines",
	}

	assert.Equal(t, expectedTodo, found[0])
}
