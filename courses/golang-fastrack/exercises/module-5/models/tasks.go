package models

type Tasks struct {
	Id int
	Title string
	Description string
	Completed bool
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}