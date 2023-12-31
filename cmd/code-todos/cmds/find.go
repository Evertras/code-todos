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
	configKeyOutput     = "output"
	configKeyLinkPrefix = "link-prefix"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find TODOs in files and directories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(cmd.ErrOrStderr(), "No files or directories specified")
			os.Exit(1)
		}

		todos, errs := todos.FindTodos(args...)

		if len(errs) > 0 {
			for file, err := range errs {
				fmt.Fprintf(cmd.ErrOrStderr(), "ERROR: %s: %s\n", file, err)
			}

			os.Exit(1)
		}

		var output string
		var err error

		switch viper.GetString(configKeyOutput) {
		case "markdown":
			markdownConfig := outputs.MarkdownTableConfig{
				LinkPrefix: viper.GetString(configKeyLinkPrefix),
			}
			output, err = outputs.MarkdownTable(todos, markdownConfig)

		case "json":
			output, err = outputs.Json(todos)
		}

		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error generating output: %s\n", err)
			os.Exit(1)
		}

		fmt.Fprintln(cmd.OutOrStdout(), output)
	},
}

func init() {
	findCmd.Flags().StringP(configKeyOutput, "o", "markdown", "Output format (markdown, json)")
	findCmd.Flags().String(configKeyLinkPrefix, "", "Link prefix for markdown output (such as .. to go up a level)")

	err := viper.BindPFlags(findCmd.Flags())

	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(findCmd)
}
