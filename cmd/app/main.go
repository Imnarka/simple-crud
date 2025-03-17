package main

import (
	"github.com/Imnarka/simple-crud/internal/database"
	"github.com/Imnarka/simple-crud/internal/handlers"
	"github.com/Imnarka/simple-crud/internal/logging"
	"github.com/Imnarka/simple-crud/internal/repositories"
	"github.com/Imnarka/simple-crud/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	database.InitDb()
	logger.InitLogger()

	logger.Info("Старт API сервера", logrus.Fields{})

	taskRepo := repositories.NewTaskRepository(database.DB)
	taskService := service.NewService(taskRepo)
	taskHandler := handlers.NewHandler(taskService)
	router := mux.NewRouter()

	router.HandleFunc("/api/task/create", taskHandler.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/task", taskHandler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/task/{id}", taskHandler.GetTaskByIDHandler).Methods("GET")
	router.HandleFunc("/api/task/{id}/update", taskHandler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/task/{id}/delete", taskHandler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
