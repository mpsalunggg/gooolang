package main

import (
	"fmt"
	"module-5/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handlers.GetTasksHandler)
	http.HandleFunc("/tasks/create", handlers.CreateTaskHandler)

	// Contoh curl untuk membuat task baru:
	// curl -X POST http://localhost:8181/tasks/create -H "Content-Type: application/json" -d '{"title":"Task Baru","description":"Deskripsi Task Baru","completed":false}'

	fmt.Println("Server berjalan di http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
