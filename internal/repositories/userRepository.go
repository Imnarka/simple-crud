package repositories

import (
	"errors"
	"fmt"
	"strings"

	repoError "github.com/Imnarka/simple-crud/internal/errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.Users) (*models.Users, error)
	GetAllUsers() ([]models.Users, error)
	GetUserById(userId uint) (*models.Users, error)
	UpdateUser(id uint, updates map[string]interface{}) (*models.Users, error)
	DeleteUser(id uint) error
	UserExists(userId uint) (bool, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (repo *userRepositoryImpl) CreateUser(user *models.Users) (*models.Users, error) {
	if err := repo.db.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, repoError.ErrUserAlreadyExists
		}
		return nil, err
	}
	return user, nil
}

func (repo *userRepositoryImpl) GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	if err := repo.db.Find(&users).Error; err != nil {
		return []models.Users{}, err
	}
	return users, nil
}

func (repo *userRepositoryImpl) GetUserById(userId uint) (*models.Users, error) {
	var user models.Users
	if err := repo.db.First(&user, "id = ?", userId).Error; err != nil {
		fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repoError.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *userRepositoryImpl) UpdateUser(id uint, updates map[string]interface{}) (*models.Users, error) {
	var user models.Users
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repoError.ErrUserNotFound
		}
		return nil, err
	}
	if err := repo.db.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepositoryImpl) DeleteUser(id uint) error {
	result := repo.db.Unscoped().Delete(&models.Users{}, id)
	if result.RowsAffected == 0 {
		return repoError.ErrUserNotFound
	}
	return result.Error
}

func (repo *userRepositoryImpl) UserExists(userId uint) (bool, error) {
	var count int64
	err := repo.db.Model(&models.Users{}).Where("id = ?", userId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
