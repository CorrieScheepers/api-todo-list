package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run the gRPC client to interact with the Todo List",
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the Todo List",
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve flags or arguments
		taskDetails, _ := cmd.Flags().GetString("task")
		if taskDetails == "" {
			fmt.Println("Please provide a task description.")
			return
		}

		// Connect to the gRPC server
		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		//client := todoservice.NewTodoServiceClient(conn)

		// Add the task using gRPC (you'd implement the actual logic in the server)
		// Example: client.AddTask(context.Background(), &todoService.AddTaskRequest{Task: taskDetails})
		fmt.Printf("Adding task: %s\n", taskDetails)
	},
}

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Get the list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to the gRPC server
		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		//client := todoservice.NewTodoServiceClient(conn)

		// Retrieve tasks using gRPC (you'd implement the actual logic in the server)
		// Example: tasks, _ := client.GetTasks(context.Background(), &todoService.GetTasksRequest{})
		fmt.Println("Fetching the list of tasks...")
	},
}

var completeCommand = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve flags or arguments
		taskID, _ := cmd.Flags().GetString("taskID")
		if taskID == "" {
			fmt.Println("Please provide a task ID.")
			return
		}

		// Connect to the gRPC server
		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		//client := todoservice.NewTodoServiceClient(conn)

		// Complete the task using gRPC (you'd implement the actual logic in the server)
		// Example: client.CompleteTask(context.Background(), &todoService.CompleteTaskRequest{TaskID: taskID})
		fmt.Printf("Marking task %s as complete...\n", taskID)
	},
}

func init() {
	clientCmd.AddCommand(addCommand)
	clientCmd.AddCommand(getCommand)
	clientCmd.AddCommand(completeCommand)

	// Add flags for each subcommand
	addCommand.Flags().String("task", "", "Task description")
	completeCommand.Flags().String("taskID", "", "Task ID")
}
