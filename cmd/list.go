package cmd

import (
	"fmt"
	"log"

	"api-todo-list/client"

	"github.com/spf13/cobra"
)

var serverAddress string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		todoClient, err := client.NewTodoClient(serverAddress)
		if err != nil {
			log.Fatalf("Could not create client: %v", err)
		}
		defer todoClient.Close()

		fmt.Println("Fetching tasks...")
		todoClient.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&serverAddress, "server", "s", "localhost:50051", "gRPC server address")
}
