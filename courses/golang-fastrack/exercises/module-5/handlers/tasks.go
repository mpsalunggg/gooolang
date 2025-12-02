package handlers

import (
	"encoding/json"
	"module-5/models"
	"net/http"
)

var tasks = []models.Tasks{
	{Id: 1, Title: "Task 1", Description: "Description 1", Completed: false},
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task := models.Tasks{
		Id:          len(tasks) + 1,
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}

	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
