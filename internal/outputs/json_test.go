package outputs

import (
	"testing"

	"github.com/evertras/code-todos/internal/todos"
)

func TestJson(t *testing.T) {
	testCases := []struct {
		name     string
		todos    []todos.Todo
		expected string
	}{
		{
			name:     "empty",
			todos:    []todos.Todo{},
			expected: "[]",
		},
		{
			name: "single todo",
			todos: []todos.Todo{
				{
					Filename:    "testdata/go/main.go",
					PackageName: "main",
					Line:        4,
					Text:        "TODO: Write a really cool thing here. And do it on multiple lines",
				},
			},
			expected: `[
  {
    "filename": "testdata/go/main.go",
    "packageName": "main",
    "line": 4,
    "text": "TODO: Write a really cool thing here. And do it on multiple lines"
  }
]`,
		},
		{
			name: "multiple todos",
			todos: []todos.Todo{
				{
					Filename:    "testdata/go/main.go",
					PackageName: "main",
					Line:        4,
					Text:        "TODO: Write a really cool thing here. And do it on multiple lines",
				},
				{
					Filename:    "testdata/go/main.go",
					PackageName: "main",
					Line:        50,
					Text:        "TODO: This is another todo",
				},
			},
			expected: `[
  {
    "filename": "testdata/go/main.go",
    "packageName": "main",
    "line": 4,
    "text": "TODO: Write a really cool thing here. And do it on multiple lines"
  },
  {
    "filename": "testdata/go/main.go",
    "packageName": "main",
    "line": 50,
    "text": "TODO: This is another todo"
  }
]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Json(tc.todos)

			if err != nil {
				t.Fatal(err)
			}

			if tc.expected != actual {
				t.Errorf("expected:\n%s\n\nactual:\n%s", tc.expected, actual)
			}
		})
	}
}
