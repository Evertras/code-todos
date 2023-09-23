package cmds

import (
	"fmt"
	"os"

	"github.com/evertras/code-todos/internal/outputs"
	"github.com/evertras/code-todos/internal/todos"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configKeyOutput = "output"
)

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

			os.Exit(1)
		}

		var output string
		var err error

		switch viper.GetString(configKeyOutput) {
		case "markdown":
			output, err = outputs.MarkdownTable(todos)

		case "json":
			output, err = outputs.Json(todos)
		}

		if err != nil {
			fmt.Printf("Error generating output: %s\n", err)
			os.Exit(1)
		}

		fmt.Println(output)
	},
}

func init() {
	findCmd.Flags().StringP(configKeyOutput, "o", "markdown", "Output format (markdown, json)")

	err := viper.BindPFlags(findCmd.Flags())

	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(findCmd)
}
