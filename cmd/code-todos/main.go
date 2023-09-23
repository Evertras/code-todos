package main

import "github.com/evertras/code-todos/cmd/code-todos/cmds"

func main() {
	err := cmds.Execute()

	if err != nil {
		panic(err)
	}
}
