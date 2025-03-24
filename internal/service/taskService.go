package service

import (
	serviceError "github.com/Imnarka/simple-crud/internal/errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/repositories"
	api "github.com/Imnarka/simple-crud/internal/web/tasks"
	"github.com/Imnarka/simple-crud/pkg/utils"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
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
	if err != nil {
		return nil, err
	}
	return task, err
}

func (s *TaskService) UpdateTask(id uint, updates *api.UpdateTask) (*models.Task, error) {
	if updates == nil {
		return nil, serviceError.ErrInvalidRequestFormat
	}
	updatesMap, err := utils.StructToMap(updates)
	if err != nil {
		return nil, err
	}
	task, err := s.repo.UpdateTaskById(id, updatesMap)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	err := s.repo.DeleteTaskById(id)
	if err != nil {
		return err
	}
	return nil
}
