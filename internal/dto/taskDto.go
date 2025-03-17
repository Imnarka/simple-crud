package dto

import "github.com/go-playground/validator/v10"

type TaskRequest struct {
	Task   string `json:"task" validate:"required"`
	IsDone *bool  `json:"is_done,omitempty"`
}

type TaskResponse struct {
	Id     uint   `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type TasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

var validate = validator.New()

func (r *TaskRequest) Validate() error {
	return validate.Struct(r)
}
