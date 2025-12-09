package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	fmt.Println("Server berjalan di http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
