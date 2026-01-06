package handlers

import (
	"encoding/json"
	"module-7/internal/storage"
	"net/http"
)

type ProductHandler struct {
	storage *storage.ProductStorage
}

func NewProductHandler(storage *storage.ProductStorage) *ProductHandler {
	return &ProductHandler{storage: storage}
}

func (ph *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		ph.handleGetProducts(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ph *ProductHandler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products := ph.storage.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
