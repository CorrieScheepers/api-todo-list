package client

import (
	"context"
	"fmt"
	"log"
	"time"

	todoservice "api-todo-list/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TodoClient struct {
	conn   *grpc.ClientConn
	client todoservice.TodoServiceClient
}

// NewTodoClient creates a new gRPC client
func NewTodoClient(serverAddr string) (*TodoClient, error) {

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}

	return &TodoClient{
		conn:   conn,
		client: todoservice.NewTodoServiceClient(conn),
	}, nil
}

// Close closes the gRPC connection
func (c *TodoClient) Close() {
	c.conn.Close()
}

// ListTasks calls the gRPC server to fetch all tasks
func (c *TodoClient) ListTasks() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := c.client.ListTasks(ctx, &todoservice.ListTasksRequest{})
	if err != nil {
		log.Fatalf("Error fetching tasks: %v", err)
	}

	fmt.Println("Tasks:")
	for _, task := range resp.Tasks {
		fmt.Printf("ID: %d, Title: %s, Completed: %v\n", task.Id, task.Title, task.Completed)
	}
}
