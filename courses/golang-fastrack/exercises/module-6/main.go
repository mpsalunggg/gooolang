package main

import (
	"fmt"
	"module-6/handlers"
	"module-6/repositories"
	"net/http"
)

func main() {
	taskRepository := repositories.NewTaskRepository()
	taskHandler := handlers.NewTaskHandler(taskRepository)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			taskHandler.GetAllTasks(w, r)
			return
		}
		if r.Method == http.MethodPost {
			taskHandler.CreateTask(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	fmt.Println("Server berjalan di http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
