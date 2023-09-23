package todos

import (
	"go/parser"
	"go/token"
	"strings"
)

func findTodosInFile(filename string) ([]Todo, error) {
	todos := make([]Todo, 0)

	fileset := token.NewFileSet()

	parsed, err := parser.ParseFile(fileset, filename, nil, parser.ParseComments)

	if err != nil {
		return nil, err
	}

	for _, comment := range parsed.Comments {
		text := comment.Text()

		if !strings.Contains(text, "TODO:") {
			continue
		}

		todos = append(todos, Todo{
			Filename:    filename,
			PackageName: parsed.Name.Name,
			Line:        fileset.Position(comment.Pos()).Line,
			Text:        strings.TrimRight(text, "\n"),
		})
	}

	return todos, nil
}
