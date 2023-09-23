package cmds

import (
	"fmt"
	"os"

	"github.com/evertras/code-todos/internal/todos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find TODOs in files and directories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No files or directories specified")
			os.Exit(1)
		}

		todos, errs := todos.FindTodos(args...)

		if len(errs) > 0 {
			fmt.Println("ERRS")

			for file, err := range errs {
				fmt.Printf("%s: %s\n", file, err)
			}
		}

		for _, todo := range todos {
			fmt.Println(todo)
		}
	},
}