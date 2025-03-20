package service

import (
	"errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/repositories"
	"gorm.io/gorm"
)

var (
	ErrTaskNotFound         = errors.New("задача не найдена")
	ErrInvalidRequestFormat = errors.New("неверный формат запроса")
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

func (s *TaskService) GetTaskById(id uint) (*models.Task, error) {
	task, err := s.repo.GetTaskById(id)
	if task == nil {
		return nil, ErrTaskNotFound
	}
	return task, err
}

func (s *TaskService) UpdateTask(id uint, updates map[string]interface{}) (*models.Task, error) {
	if len(updates) == 0 {
		return nil, ErrInvalidRequestFormat
	}
	task, err := s.repo.UpdateTaskById(id, updates)
	if err != nil {
		return nil, err
	}
	return task, err
}

func (s *TaskService) DeleteTask(id uint) error {
	err := s.repo.DeleteTaskById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrTaskNotFound
	}
	return err
}
