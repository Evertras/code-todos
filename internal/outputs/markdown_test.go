package outputs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/evertras/code-todos/internal/todos"
)

func TestMarkdownTable(t *testing.T) {
	testCases := []struct {
		name     string
		todos    []todos.Todo
		config   MarkdownTableConfig
		expected string
	}{
		{
			name:   "empty",
			todos:  []todos.Todo{},
			config: MarkdownTableConfig{},
			expected: `| Filename | Package | Line | Text |
| -------- | ------- | ---- | ---- |
`,
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
			config: MarkdownTableConfig{},
			expected: `| Filename | Package | Line | Text |
| -------- | ------- | ---- | ---- |
| [testdata/go/main.go](testdata/go/main.go#L4) | main | 4 | TODO: Write a really cool thing here. And do it on multiple lines |
`,
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
			config: MarkdownTableConfig{},
			expected: `| Filename | Package | Line | Text |
| -------- | ------- | ---- | ---- |
| [testdata/go/main.go](testdata/go/main.go#L4) | main | 4 | TODO: Write a really cool thing here. And do it on multiple lines |
| [testdata/go/main.go](testdata/go/main.go#L50) | main | 50 | TODO: This is another todo |
`,
		},
		{
			name: "with link prefix",
			todos: []todos.Todo{
				{
					Filename:    "testdata/go/main.go",
					PackageName: "main",
					Line:        4,
					Text:        "TODO: Write a really cool thing here. And do it on multiple lines",
				},
			},
			config: MarkdownTableConfig{
				LinkPrefix: "../",
			},
			expected: `| Filename | Package | Line | Text |
| -------- | ------- | ---- | ---- |
| [testdata/go/main.go](../testdata/go/main.go#L4) | main | 4 | TODO: Write a really cool thing here. And do it on multiple lines |
`,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual, err := MarkdownTable(test.todos, test.config)

			if err != nil {
				t.Fatal(err)
			}

			if actual != test.expected {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}
