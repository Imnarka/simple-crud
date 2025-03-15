package handlers

import (
	"encoding/json"
	"github.com/Imnarka/simple-crud.git/internal/database"
	"github.com/Imnarka/simple-crud.git/internal/dto"
	logger "github.com/Imnarka/simple-crud.git/internal/logging"
	"github.com/Imnarka/simple-crud.git/internal/models"
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
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Задача создана"))
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
