package main

import (
	"os"

	"github.com/evertras/code-todos/cmd/code-todos/cmds"
)

func main() {
	err := cmds.Execute()

	if err != nil {
		os.Exit(1)
	}
}
