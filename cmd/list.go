package cmd

import (
	"fmt"
	"log"

	"api-todo-list/client"

	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load(".config/.config")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := os.Getenv("SERVER")
	listCmd.Flags().StringVarP(&serverAddress, "server", "s", server, "gRPC server address")
}
