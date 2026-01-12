package main

import (
	"fmt"
	"module-7/internal/handlers"
	"module-7/internal/storage"
	"net/http"
)

func main() {
	productStorage := storage.NewProductStorage()
	productHandler := handlers.NewProductHandler(productStorage)

	http.HandleFunc("/products", productHandler.HandleProducts)
	http.HandleFunc("/products/", handleProductRoutes(productHandler))

	fmt.Println("Server berjalan di http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}


func handleProductRoutes(ph *handlers.ProductHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// path := r.URL.Path

		ph.HandleProducts(w, r)
	}
}