package repositories

import (
	"errors"

	repoError "github.com/Imnarka/simple-crud/internal/errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id uint) (*models.Task, error)
	UpdateTaskById(id uint, updates map[string]interface{}) (*models.Task, error)
	DeleteTaskById(id uint) error
	GetTasksByUserId(userId uint) ([]models.Task, error)
}

// taskRepository структура репозиторя
type taskRepositoryImpl struct {
	db *gorm.DB
}

// NewTaskRepository Конструктор taskRepository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryImpl{db: db}
}

// CreateTask создане задачи
func (r *taskRepositoryImpl) CreateTask(task *models.Task) (*models.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// GetAllTasks получение всех задач из БД
func (r *taskRepositoryImpl) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return []models.Task{}, nil
	}
	return tasks, err
}

// GetTaskById получение задачи по id
func (r *taskRepositoryImpl) GetTaskById(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &task, err
}

// UpdateTaskById обновление задачи по ID
func (r *taskRepositoryImpl) UpdateTaskById(id uint, updates map[string]interface{}) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repoError.ErrUserNotFound
		}
		return nil, err
	}
	if err := r.db.Model(&task).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTaskById удаление задачи по ID
func (r *taskRepositoryImpl) DeleteTaskById(id uint) error {
	result := r.db.Unscoped().Delete(&models.Task{}, id)
	if result.RowsAffected == 0 {
		return repoError.ErrUserNotFound
	}
	return result.Error
}

// GetTasksByUserId получение задачи по ID пользователя
func (r *taskRepositoryImpl) GetTasksByUserId(userId uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userId).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
