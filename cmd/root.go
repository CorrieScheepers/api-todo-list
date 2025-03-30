package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Root command of the CLI
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "CLI for managing tasks via gRPC",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand: list, add, etc.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}

// Initialize commands
func init() {
	rootCmd.AddCommand(listCmd)
}
