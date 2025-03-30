package server

import (
	"context"

	todoservice "api-todo-list/proto"
	"api-todo-list/repository"
)

type TodoServiceServer struct {
	todoservice.UnimplementedTodoServiceServer
	TaskRepo *repository.TaskRepository
}

func NewServer(repo *repository.TaskRepository) *TodoServiceServer {
	return &TodoServiceServer{TaskRepo: repo}
}

func (s *TodoServiceServer) ListTasks(ctx context.Context, req *todoservice.ListTasksRequest) (*todoservice.ListTasksResponse, error) {
	tasks, err := s.TaskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	grpcTasks := []*todoservice.Task{}
	for _, task := range tasks {
		grpcTasks = append(grpcTasks, &todoservice.Task{
			Id:          int32(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		})
	}

	return &todoservice.ListTasksResponse{Tasks: grpcTasks}, nil
}
