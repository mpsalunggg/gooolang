package main

import (
	"fmt"
	"module-5/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handlers.GetTasksHandler)

	fmt.Println("Server berjalan di http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
