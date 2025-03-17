package handlers

import (
	"encoding/json"
	"errors"
	"github.com/Imnarka/simple-crud/internal/dto"
	logger "github.com/Imnarka/simple-crud/internal/logging"
	"github.com/Imnarka/simple-crud/internal/models"
	"github.com/Imnarka/simple-crud/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *service.TaskService
}

func NewHandler(service *service.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		msg := "Неверный формат запроса"
		logger.Error(msg, logrus.Fields{"error": err})
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		msg := "Ошибка валидаци тела запроса"
		logger.Error(msg, logrus.Fields{"error": err})
		http.Error(w, msg, http.StatusUnprocessableEntity)
		return
	}
	task := models.Task{
		Task:   req.Task,
		IsDone: false,
	}

	createdTask, err := h.Service.CreateTask(&task)
	if err != nil {
		msg := "Ошибка при создании задачи"
		logger.Error(msg, logrus.Fields{"error": err})
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	response := dto.TaskResponse{
		Id:     createdTask.ID,
		Task:   createdTask.Task,
		IsDone: createdTask.IsDone,
	}

	logger.Info("Задача успешно создана", logrus.Fields{"task_id": createdTask.ID})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса /task/create", logrus.Fields{})
	tasks, err := h.Service.GetTasks()
	if err != nil {
		msg := "Ошибка получения данных в БД"
		logger.Error(msg, logrus.Fields{"error": err})
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	response := dto.TasksResponse{Tasks: []dto.TaskResponse{}}
	for _, task := range tasks {
		response.Tasks = append(response.Tasks, dto.TaskResponse{
			Id:     task.ID,
			Task:   task.Task,
			IsDone: task.IsDone,
		})
	}
	logger.Info("Задачи успешно получены: ", logrus.Fields{"tasks": len(response.Tasks)})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Tasks)
}

func (h *Handler) GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса GET /task/{id}", logrus.Fields{})

	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "Некорректный ID задачи"
		logger.Error(msg, logrus.Fields{"error": err, "task_id": vars["id"]})
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	task, err := h.Service.GetTaskById(uint(taskId))
	switch {
	case errors.Is(err, service.ErrTaskNotFound):
		http.Error(w, "Задача не найдена", http.StatusNotFound)
	case err != nil:
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
	default:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.TaskResponse{
			Id: task.ID, Task: task.Task, IsDone: task.IsDone,
		})
	}
}

func (h *Handler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса PATCH /task/{id}/update", logrus.Fields{})
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		msg := "Некорректный ID задачи"
		logger.Error(msg, logrus.Fields{"error": err, "task_id": vars["id"]})
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		msg := "Неверный формат запроса"
		logger.Error(msg, logrus.Fields{"error": err})
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	task, err := h.Service.UpdateTask(uint(taskId), updates)
	switch {
	case errors.Is(err, service.ErrTaskNotFound):
		http.Error(w, "Задача не найдена", http.StatusNotFound)
	case err != nil:
		http.Error(w, "Ошибка обновления задачи", http.StatusInternalServerError)
	default:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dto.TaskResponse{
			Id: task.ID, Task: task.Task, IsDone: task.IsDone,
		})
	}
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса DELETE /task/{id}/delete", logrus.Fields{})
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "Некорректный ID задачи"
		logger.Error(msg, logrus.Fields{"error": err, "task_id": vars["id"]})
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteTask(uint(taskId))
	switch {
	case errors.Is(err, service.ErrTaskNotFound):
		http.Error(w, "Задача не найдена", http.StatusNotFound)
	case err != nil:
		http.Error(w, "Ошибка удаления задачи", http.StatusInternalServerError)
	default:
		logger.Info("Задача успешно удалена", logrus.Fields{"task_id": taskId})
		w.WriteHeader(http.StatusNoContent)
	}
}
