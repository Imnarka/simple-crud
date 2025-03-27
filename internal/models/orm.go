package models

import "gorm.io/gorm"

// Task Схема таблицы task
type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
}

// Users схема таблицы users
type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Task     []Task `json:"tasks" gorm:"foreignKey:UserID"`
}
