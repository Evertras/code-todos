package todos

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindTodos(paths ...string) ([]Todo, map[string]error) {
	errs := map[string]error{}
	todos := make([]Todo, 0)

	for _, path := range paths {
		pathTodos, err := findTodosInPath(path)

		if err != nil {
			errs[path] = err
			continue
		}

		todos = append(todos, pathTodos...)
	}

	return todos, errs
}

// TODO: add globbing
func findTodosInPath(path string) ([]Todo, error) {
	todos := make([]Todo, 0)

	// We use filepath.WalkDir instead of parser.ParseDir because we want to
	// better track file names and ignore files in gitignore, etc.
	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		info, err := d.Info()

		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %w", path, err)
		}

		// TODO: Filter better
		if filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		pathTodos, err := findTodosInFile(path)

		if err != nil {
			// TODO: Better error handling
			return err
		}

		todos = append(todos, pathTodos...)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return todos, nil
}
