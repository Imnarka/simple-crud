package repositories

import (
	"github.com/Imnarka/simple-crud/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id uint) (*models.Task, error)
	UpdateTaskById(id uint, updates map[string]interface{}) (*models.Task, error)
	DeleteTaskById(id uint) error
}

// taskRepository структура репозиторя
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository Конструктор taskRepository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// CreateTask создане задачи
func (r *taskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// GetAllTasks получение всех задач из БД
func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return []models.Task{}, nil
	}
	return tasks, err
}

func (r *taskRepository) GetTaskById(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, nil
	}
	return &task, err
}

// UpdateTaskById обновление задачи по ID
func (r *taskRepository) UpdateTaskById(id uint, updates map[string]interface{}) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.Model(&task).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTaskById удаление задачи по ID
func (r *taskRepository) DeleteTaskById(id uint) error {
	result := r.db.Unscoped().Delete(&models.Task{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
