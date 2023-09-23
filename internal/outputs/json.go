package outputs

import (
	"encoding/json"

	"github.com/evertras/code-todos/internal/todos"
)

func Json(todos []todos.Todo) (string, error) {
	marshaled, err := json.MarshalIndent(todos, "", "  ")

	if err != nil {
		return "", err
	}

	return string(marshaled), nil
}
