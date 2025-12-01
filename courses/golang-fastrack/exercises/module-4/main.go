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

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte("OK"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound) // 404
	w.Write([]byte("Not Found"))
}

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest) // 400
	w.Write([]byte("Bad Request"))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte("Selamat datang di API"))
	nama := r.URL.Query().Get("nama")
	w.Write([]byte("Hello bapak: " + nama))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/ok", okHandler)
	http.HandleFunc("/notfound", notFoundHandler)
	http.HandleFunc("/badrequest", badRequestHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	fmt.Println("Server berjalan di http://localhost:8181")

	http.ListenAndServe(":8181", nil)
}
