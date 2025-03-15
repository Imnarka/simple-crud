package main

import (
	"github.com/Imnarka/simple-crud.git/internal/database"
	"github.com/Imnarka/simple-crud.git/internal/handlers"
	"github.com/Imnarka/simple-crud.git/internal/logging"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	database.InitDb()
	logger.InitLogger()
	logger.Info("Старт API сервера", logrus.Fields{})
	router := mux.NewRouter()
	router.HandleFunc("/api/task/create", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/task", handlers.GetTaskHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
