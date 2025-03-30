package server

import (
	"api-todo-list/grpc"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	// Adjust the import path to your generated grpc package
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

// server struct holds the db connection for grpc methods
type server struct {
	grpc.UnimplementedTodoServiceServer
	db *sql.DB
}

// CreateTask inserts a new task into the database
func (s *server) CreateTask(ctx context.Context, req *grpc.CreateTaskRequest) (*grpc.CreateTaskResponse, error) {
	// Insert task into the database
	query := "INSERT INTO tasks (title, description) VALUES (?, ?)"
	result, err := s.db.Exec(query, req.GetTitle(), req.GetDescription())
	if err != nil {
		return nil, fmt.Errorf("failed to insert task: %v", err)
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %v", err)
	}

	// Create the task response with ID and other details
	task := &grpc.Task{
		Id:          int32(taskID),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Completed:   false, // Default to incomplete
	}

	return &grpc.CreateTaskResponse{Task: task}, nil
}

func main() {
	// Connect to the database
	db, err := database.connectToMySqlDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Ensure the connection is closed when the application ends

	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	todoService := &server{db: db}

	// Register the TodoService server with gRPC
	grpc.RegisterTodoServiceServer(grpcServer, todoService)

	// Start the gRPC server and listen on the specified port
	fmt.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
