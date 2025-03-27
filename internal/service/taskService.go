package service

import (
	serviceError "github.com/Imnarka/simple-crud/internal/errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/repositories"
	api "github.com/Imnarka/simple-crud/internal/web/tasks"
	"github.com/Imnarka/simple-crud/pkg/utils"
)

type TaskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	return s.taskRepo.CreateTask(task)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.taskRepo.GetAllTasks()
}

func (s *TaskService) GetTaskById(id uint) (*models.Task, error) {
	task, err := s.taskRepo.GetTaskById(id)
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
	task, err := s.taskRepo.UpdateTaskById(id, updatesMap)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id uint) error {
	err := s.taskRepo.DeleteTaskById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetTasksByUserId(userId uint) ([]models.Task, error) {
	tasks, err := s.taskRepo.GetTasksByUserId(userId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
