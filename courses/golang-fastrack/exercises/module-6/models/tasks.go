package models

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
