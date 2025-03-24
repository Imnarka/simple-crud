package service

import (
	serviceError "github.com/Imnarka/simple-crud/internal/errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/repositories"
	api "github.com/Imnarka/simple-crud/internal/web/users"
	"github.com/Imnarka/simple-crud/pkg/utils"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.Users) (*models.Users, error) {
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) GetAllUsers() ([]models.Users, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserById(userId uint) (*models.Users, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(id uint, updates *api.UpdateUser) (*models.Users, error) {
	if updates == nil {
		return nil, serviceError.ErrInvalidRequestFormat
	}
	userUpdates, err := utils.StructToMap(updates)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.UpdateUser(id, userUpdates)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id uint) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
