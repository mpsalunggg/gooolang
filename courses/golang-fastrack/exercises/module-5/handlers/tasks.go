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
