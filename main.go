package main

import (
	"api-todo-list/database"
	todoservice "api-todo-list/grpc"
	"api-todo-list/repository"
	"api-todo-list/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := database.ConnectToMySqlDb()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	taskRepo := repository.NewTaskRepository(db)
	grpcServer := grpc.NewServer()
	todoServiceServer := server.NewServer(taskRepo)
	todoservice.RegisterTodoServiceServer(grpcServer, todoServiceServer)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	fmt.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
