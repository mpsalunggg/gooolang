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

	http.HandleFunc("/tasks", taskHandler.HandleTasks)
	http.HandleFunc("/tasks/", taskHandler.HandleTaskByID)

	fmt.Println("Server berjalan di http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
