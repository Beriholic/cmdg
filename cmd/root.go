package cmd

import (
	"fmt"
	"os"

	"github.com/Beriholic/cmdg/internal"
	"github.com/spf13/cobra"
)

var userCommond = ""

func init() {
	rootCmd.Flags().StringVarP(&userCommond, "command", "c", "", "command to execute")
}

var rootCmd = &cobra.Command{
	Use:   "cmdg -c <command>",
	Short: "A CLI tool to get terminal commands from natural language",
	Long:  `cmdg is a command-line interface tool that converts natural language descriptions into terminal commands, making it easier to find and use terminal commands without memorizing them`,
	Example: `  # Get command to list all files in current directory
  cmdg "show all files in current directory"

  # Get command to create a new directory
  cmdg -c "create a new folder named test"

  # Get command to find large files
  cmdg -c "find files larger than 100MB"`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		err := internal.GeneratorCommand(ctx, userCommond)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
