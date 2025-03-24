package handlers

import (
	"context"
	"errors"

	serviceError "github.com/Imnarka/simple-crud/internal/errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/service"
	api "github.com/Imnarka/simple-crud/internal/web/users"
	openapitypes "github.com/oapi-codegen/runtime/types"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(context.Context, api.GetUsersRequestObject) (api.GetUsersResponseObject, error) {
	userList, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := make(api.GetUsers200JSONResponse, 0, len(userList))
	for _, user := range userList {
		email := openapitypes.Email(user.Email)
		response = append(response, api.UsersResponse{
			Id:    user.ID,
			Email: email,
		})
	}
	return response, err
}

func (h *UserHandler) CreateUser(_ context.Context, request api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	userRequest := request.Body
	if userRequest.Email == "" || userRequest.Password == "" {
		return api.CreateUser400Response{}, nil
	}
	userCreate := models.Users{
		Email:    string(userRequest.Email),
		Password: userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(&userCreate)
	if err != nil {
		if errors.Is(err, serviceError.ErrUserAlreadyExists) {
			return api.CreateUser422Response{}, nil
		}
		return api.CreateUser500Response{}, err
	}
	response := api.CreateUser200JSONResponse{
		Id:    createdUser.ID,
		Email: openapitypes.Email(createdUser.Email),
	}
	return response, nil
}

func (h *UserHandler) UpdateUser(_ context.Context, request api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error) {
	updatedUser, err := h.Service.UpdateUser(request.Id, request.Body)
	if err != nil {
		return api.UpdateUser500Response{}, err
	}
	response := api.UpdateUser200JSONResponse{
		Id:    updatedUser.ID,
		Email: (openapitypes.Email)(updatedUser.Email),
	}
	return response, nil
}

func (h *UserHandler) GetUserById(ctx context.Context, request api.GetUserByIdRequestObject) (api.GetUserByIdResponseObject, error) {
	user, err := h.Service.GetUserById(request.Id)
	if err != nil {
		if errors.Is(err, serviceError.ErrUserNotFound) {
			return api.GetUserById404Response{}, nil
		}
		return api.GetUserById500Response{}, nil
	}
	response := api.GetUserById200JSONResponse{
		Id:    user.ID,
		Email: (openapitypes.Email)(user.Email),
	}
	return response, nil
}

func (h *UserHandler) DeleteUserById(ctx context.Context, request api.DeleteUserByIdRequestObject) (api.DeleteUserByIdResponseObject, error) {
	err := h.Service.DeleteUser(request.Id)
	if err != nil {
		if errors.Is(err, serviceError.ErrUserNotFound) {
			return api.DeleteUserById404Response{}, nil
		}
		return api.DeleteUserById500Response{}, err
	}
	return api.DeleteUserById204Response{}, nil
}
