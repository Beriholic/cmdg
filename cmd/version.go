package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const verison = "0.1.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of the cmdg",
	Long:  `print the version of the cmdg`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cmdg version %s\n", verison)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
