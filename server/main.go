package main

import (
    "context"
    "fmt"
    "log"

    "github.com/CorrieScheepers2/api-todo-list/grpc"
    "google.golang.org/grpc"
    "database/sql"
)

type server struct {
    grpc.UnimplementedTodoServiceServer
    db *sql.DB
}

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
    db, err := connectToDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Set up gRPC server
    grpcServer := grpc.NewServer()
    todoService := &server{db: db}

    grpc.RegisterTodoServiceServer(grpcServer, todoService)

    // Start listening on a port (e.g., 50051)
    fmt.Println("Server is listening on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}
