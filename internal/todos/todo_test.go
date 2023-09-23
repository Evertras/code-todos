package todos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTodosInFile(t *testing.T) {
	todos, err := findTodosInFile("testdata/go/main.go")

	assert.NoError(t, err)
	assert.Len(t, todos, 1)

	expectedTodo := Todo{
		Filename:    "testdata/go/main.go",
		PackageName: "main",
		Line:        4,
		Text:        "TODO: Write a really cool thing here.\nAnd do it on multiple lines",
	}

	assert.Equal(t, expectedTodo, todos[0])
}
