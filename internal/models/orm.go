package models

import "gorm.io/gorm"

// Task Схема таблицы task
type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

// Users схема таблицы users
type Users struct {
	gorm.Model
	Email    string `json:"task" gorm:"unique"`
	Password string `json:"password"`
}
