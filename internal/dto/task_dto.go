package dto

type TaskRequestBody struct {
	Task   string `json:"task,omitempty"`
	IsDone *bool  `json:"is_done,omitempty"`
}
