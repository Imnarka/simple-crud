package handlers

import (
	"context"
	"errors"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/service"
	"github.com/Imnarka/simple-crud/internal/web/tasks"
)

type Handler struct {
	Service *service.TaskService
}

func NewHandler(service *service.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTasks(context.Context, tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	taskList, err := h.Service.GetTasks()
	if err != nil {
		return nil, err
	}
	response := make(tasks.GetTasks200JSONResponse, 0, len(taskList))
	for _, task := range taskList {
		response = append(response, tasks.Task{
			Id:     &task.ID,
			IsDone: &task.IsDone,
			Task:   task.Task,
		})
	}
	return response, nil
}

func (h *Handler) GetTaskByID(_ context.Context, request tasks.GetTaskByIDRequestObject) (tasks.GetTaskByIDResponseObject, error) {
	task, err := h.Service.GetTaskById(request.Id)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			return tasks.GetTaskByID404Response{}, nil
		}
		return tasks.GetTaskByID500Response{}, nil
	}
	return tasks.GetTaskByID200JSONResponse{
		Id:     &task.ID,
		IsDone: &task.IsDone,
		Task:   task.Task,
	}, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	if request.Body.Task == "" {
		return tasks.PostTasks422Response{}, nil
	}
	taskToCreate := models.Task{
		Task:   taskRequest.Task,
		IsDone: false,
	}
	createdTask, err := h.Service.CreateTask(&taskToCreate)
	if err != nil {
		return tasks.PostTasks500Response{}, err
	}
	response := tasks.PostTasks200JSONResponse{
		Id:     &createdTask.ID,
		Task:   createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) UpdateTask(_ context.Context, request tasks.UpdateTaskRequestObject) (tasks.UpdateTaskResponseObject, error) {
	updatedTask, err := h.Service.UpdateTask(request.Id, *request.Body)
	if err != nil {
		if errors.Is(err, service.ErrInvalidRequestFormat) {
			return tasks.UpdateTask400Response{}, nil
		}
		if errors.Is(err, service.ErrTaskNotFound) {
			return tasks.UpdateTask404Response{}, nil
		}
		return tasks.UpdateTask500Response{}, nil
	}
	return tasks.UpdateTask200JSONResponse{
		Id:     &updatedTask.ID,
		IsDone: &updatedTask.IsDone,
		Task:   updatedTask.Task,
	}, nil
}

func (h *Handler) DeleteTask(_ context.Context, request tasks.DeleteTaskRequestObject) (tasks.DeleteTaskResponseObject, error) {
	err := h.Service.DeleteTask(request.Id)
	if err != nil {
		if errors.Is(err, service.ErrTaskNotFound) {
			return tasks.DeleteTask404Response{}, nil
		}
		return tasks.DeleteTask500Response{}, nil
	}
	return tasks.DeleteTask204Response{}, nil
}
