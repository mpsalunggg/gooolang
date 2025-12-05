package handlers

import (
	"encoding/json"
	"module-6/models"
	"module-6/repositories"
	"net/http"
)

type TaskHandler struct {
	repository *repositories.TaskRepository
}

func NewTaskHandler(repository *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{repository: repository}
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.repository.GetAllTasks())
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.repository.CreateTask(task)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}
