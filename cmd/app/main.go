package main

import (
	"github.com/Imnarka/simple-crud/internal/database"
	"github.com/Imnarka/simple-crud/internal/handlers"
	"github.com/Imnarka/simple-crud/internal/logging"
	"github.com/Imnarka/simple-crud/internal/repositories"
	"github.com/Imnarka/simple-crud/internal/service"
	tasks "github.com/Imnarka/simple-crud/internal/web/tasks"
	users "github.com/Imnarka/simple-crud/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	database.InitDb()
	logger.InitLogger()

	logger.Info("Старт API сервера", logrus.Fields{})
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskRepo := repositories.NewTaskRepository(database.DB)
	newTaskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(newTaskService)
	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	userRepo := repositories.NewUserRepository(database.DB)
	newUserService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(newUserService)
	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	if err := e.Start(":8080"); err != nil {
		logger.Error("failed to start server", logrus.Fields{"err": err})
	}
}
