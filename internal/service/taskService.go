package service

import (
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/repositories"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskById(id uint) (*models.Task, error) { return s.repo.GetTaskById(id) }

func (s *TaskService) UpdateTask(id uint, updates map[string]interface{}) (*models.Task, error) {
	return s.repo.UpdateTaskById(id, updates)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskById(id)
}
