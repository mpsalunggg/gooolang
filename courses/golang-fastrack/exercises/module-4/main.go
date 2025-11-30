package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s\n", r.Method)

	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)

	fmt.Fprintf(w, "Query String: %s\n", r.URL.RawQuery)

	nama := r.URL.Query().Get("nama")
	fmt.Fprintf(w, "Nama: %s\n", nama)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/test", testHandler)

	fmt.Println("Server berjalan di http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
