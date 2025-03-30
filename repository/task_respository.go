package repository

import (
	"api-todo-list/models"

	"github.com/jinzhu/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAllTasks() ([]models.TaskModel, error) {
	var tasks []models.TaskModel

	// Query the database using the injected DB instance
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
