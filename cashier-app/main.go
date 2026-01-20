package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var product = []Product{
	{ID: 1, Name: "Product 1", Price: 10000, Stock: 10},
	{ID: 2, Name: "Product 2", Price: 20000, Stock: 20},
	{ID: 3, Name: "Product 3", Price: 30000, Stock: 30},
}

func main() {
	// for detail
	http.HandleFunc("/api/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			id := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
			for i, p := range product {
				if strconv.Itoa(product[i].ID) == id {
					json.NewEncoder(w).Encode(p)
					return
				}
			}
			http.Error(w, "product not found", http.StatusNotFound)
			return
		case "PUT":
			w.Header().Set("Content-Type", "application/json")
			id := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
			idStr, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "invalid id", http.StatusBadRequest)
				return
			}
			var updatedProduct Product
			err = json.NewDecoder(r.Body).Decode(&updatedProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			for i, p := range product {
				if p.ID == idStr {
					updatedProduct.ID = p.ID
					product[i] = updatedProduct
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(product[i])
					return
				}
			}
			http.Error(w, "product not found", http.StatusNotFound)
			return
		case "DELETE":
			w.Header().Set("Content-Type", "application/json")
			id := strings.TrimPrefix(r.URL.Path, "/api/v1/products/")
			for i, p := range product {
				if strconv.Itoa(p.ID) == id {
					product = append(product[:i], product[i+1:]...)
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(product)
					return
				}
			}
			http.Error(w, "product not found", http.StatusNotFound)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		case "POST":
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			product = append(product, newProduct)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newProduct)
			return
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
