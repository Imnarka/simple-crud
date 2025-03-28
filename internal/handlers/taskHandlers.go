package handlers

import (
	"context"
	"errors"
	serviceError "github.com/Imnarka/simple-crud/internal/errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/service"
	api "github.com/Imnarka/simple-crud/internal/web/tasks"
)

type TaskHandler struct {
	Service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: service,
	}
}

func (h *TaskHandler) GetTasks(context.Context, api.GetTasksRequestObject) (api.GetTasksResponseObject, error) {
	taskList, err := h.Service.GetTasks()
	if err != nil {
		return nil, err
	}
	response := make(api.GetTasks200JSONResponse, 0, len(taskList))
	for _, task := range taskList {
		response = append(response, api.Task{
			Id:     task.ID,
			IsDone: &task.IsDone,
			Task:   task.Task,
		})
	}
	return response, nil
}

func (h *TaskHandler) GetTaskByID(_ context.Context, request api.GetTaskByIDRequestObject) (api.GetTaskByIDResponseObject, error) {
	task, err := h.Service.GetTaskById(request.Id)
	if err != nil {
		if errors.Is(err, serviceError.ErrTaskNotFound) {
			return api.GetTaskByID404Response{}, nil
		}
		return api.GetTaskByID500Response{}, nil
	}
	return api.GetTaskByID200JSONResponse{
		Id:     task.ID,
		IsDone: &task.IsDone,
		Task:   task.Task,
	}, nil
}

func (h *TaskHandler) CreateTask(_ context.Context, request api.CreateTaskRequestObject) (api.CreateTaskResponseObject, error) {
	taskRequest := request.Body
	if request.Body.Task == "" {
		return api.CreateTask422Response{}, nil
	}
	taskToCreate := models.Task{
		Task:   taskRequest.Task,
		IsDone: false,
		UserID: taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(&taskToCreate)
	if err != nil {
		return api.CreateTask500Response{}, err
	}
	response := api.CreateTask200JSONResponse{
		Id:     createdTask.ID,
		Task:   createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) UpdateTask(_ context.Context, request api.UpdateTaskRequestObject) (api.UpdateTaskResponseObject, error) {
	updatedTask, err := h.Service.UpdateTask(request.Id, request.Body)
	if err != nil {
		if errors.Is(err, serviceError.ErrInvalidRequestFormat) {
			return api.UpdateTask400Response{}, nil
		}
		if errors.Is(err, serviceError.ErrTaskNotFound) {
			return api.UpdateTask404Response{}, nil
		}
		return api.UpdateTask500Response{}, nil
	}
	return api.UpdateTask200JSONResponse{
		Id:     updatedTask.ID,
		IsDone: &updatedTask.IsDone,
		Task:   updatedTask.Task,
	}, nil
}

func (h *TaskHandler) DeleteTask(_ context.Context, request api.DeleteTaskRequestObject) (api.DeleteTaskResponseObject, error) {
	err := h.Service.DeleteTask(request.Id)
	if err != nil {
		if errors.Is(err, serviceError.ErrTaskNotFound) {
			return api.DeleteTask404Response{}, nil
		}
		return api.DeleteTask500Response{}, nil
	}
	return api.DeleteTask204Response{}, nil
}

func (h *TaskHandler) GetTasksByUserID(_ context.Context, request api.GetTasksByUserIDRequestObject) (api.GetTasksByUserIDResponseObject, error) {
	taskList, err := h.Service.GetTasksByUserId(request.UserId)
	if err != nil {
		return api.GetTasksByUserID500Response{}, nil
	}
	response := make(api.GetTasksByUserID200JSONResponse, 0, len(taskList))
	for _, task := range taskList {
		response = append(response, api.Task{
			Id:     task.ID,
			IsDone: &task.IsDone,
			Task:   task.Task,
		})
	}
	return response, nil
}
