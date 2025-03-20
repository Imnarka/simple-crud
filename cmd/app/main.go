package main

import (
	"github.com/Imnarka/simple-crud/internal/database"
	"github.com/Imnarka/simple-crud/internal/handlers"
	"github.com/Imnarka/simple-crud/internal/logging"
	"github.com/Imnarka/simple-crud/internal/repositories"
	"github.com/Imnarka/simple-crud/internal/service"
	"github.com/Imnarka/simple-crud/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	database.InitDb()
	logger.InitLogger()

	logger.Info("Старт API сервера", logrus.Fields{})
	repo := repositories.NewTaskRepository(database.DB)
	newService := service.NewService(repo)
	handler := handlers.NewHandler(newService)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		logger.Error("failed to start server", logrus.Fields{"err": err})
	}
}
