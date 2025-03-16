package handlers

import (
	"encoding/json"
	"github.com/Imnarka/simple-crud.git/internal/database"
	"github.com/Imnarka/simple-crud.git/internal/dto"
	logger "github.com/Imnarka/simple-crud.git/internal/logging"
	"github.com/Imnarka/simple-crud.git/internal/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса /task/create", logrus.Fields{})
	var req dto.TaskRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	task := models.Task{Task: req.Task, IsDone: false}
	result := database.DB.Create(&task)
	if result.Error != nil {
		logger.Error("Ошибка сохранения данных в БД", logrus.Fields{"error": result.Error})
		http.Error(w, "Ошибка сохранения в БД", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(task); err != nil {
		logger.Error("Ошибка кодирования ответа", logrus.Fields{"error": err})
		http.Error(w, "Ошибка формирования ответа", http.StatusInternalServerError)
	}
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса /task/create", logrus.Fields{})
	var tasks []models.Task
	result := database.DB.Find(&tasks)
	if result.Error != nil {
		logger.Error("Ошибка сохранения данных в БД", logrus.Fields{"error": result.Error})
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса /task/{id}/update", logrus.Fields{})
	vars := mux.Vars(r)
	taskId := vars["id"]
	var req dto.TaskRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var task models.Task
	result := database.DB.First(&task, taskId)
	if result.Error != nil {
		logger.Error("Ошибка поиска задачи", logrus.Fields{"error": result.Error})
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}
	updates := make(map[string]interface{})
	if req.Task != "" {
		updates["task"] = req.Task
	}
	if req.IsDone != nil {
		updates["is_done"] = req.IsDone
	}
	if len(updates) > 0 {
		result := database.DB.Model(&task).Updates(updates)
		if result.Error != nil {
			logger.Error("Ошибка обновления задачи", logrus.Fields{"error": result.Error})
			http.Error(w, "Ошибка обновления в БД", http.StatusInternalServerError)
			return
		}
		logger.Info("Задача обновлена", logrus.Fields{"updates": updates})
	} else {
		logger.Info("Нет полей для обновления", logrus.Fields{})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запооса /task/delete", logrus.Fields{})
	vars := mux.Vars(r)
	taskId := vars["id"]
	var task models.Task
	result := database.DB.First(&task, taskId)
	if result.Error != nil {
		logger.Error("Задача не найдена", logrus.Fields{"error": result.Error})
		http.Error(w, "Задача не найдена", http.StatusNotFound)
	}
	result = database.DB.Unscoped().Delete(&task)
	if result.Error != nil {
		logger.Error("Ошибка удаления задачи", logrus.Fields{"error": result.Error})
		http.Error(w, "Ошибка удаления из БД", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Выполнение запроса /task/{id}", logrus.Fields{})

	vars := mux.Vars(r)
	taskID := vars["id"]

	var task models.Task
	result := database.DB.First(&task, taskID)
	if result.Error != nil {
		logger.Error("Ошибка поиска задачи", logrus.Fields{"error": result.Error})
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
