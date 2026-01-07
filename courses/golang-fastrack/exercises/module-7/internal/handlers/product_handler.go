package handlers

import (
	"encoding/json"
	"module-7/internal/storage"
	"net/http"
	"module-7/internal/models"
	"module-7/internal/utils"
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
	case http.MethodPost:
		ph.handleCreateProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ph *ProductHandler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products := ph.storage.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (ph *ProductHandler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = utils.ValidateProduct(req.Name, req.Price, req.Stock, req.SKU)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product := ph.storage.Create(req)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}