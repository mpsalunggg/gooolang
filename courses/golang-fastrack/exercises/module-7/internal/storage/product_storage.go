package storage

import (
	"module-7/internal/models"
	"time"
)

type ProductStorage struct {
	products map[int]models.Product
	nextID   int
}

func NewProductStorage() *ProductStorage {
	ps := &ProductStorage{
		products: make(map[int]models.Product),
		nextID:   1,
	}

	defaultProduct := models.Product{
		ID:          ps.nextID,
		Name:        "Laptop Gaming",
		Description: "High performance gaming laptop",
		Price:       15000000.00,
		Stock:       10,
		SKU:         "LAPTOP-001",
		Category:    "Electronics",
		CreatedAt:   time.Now(),
	}

	ps.products[defaultProduct.ID] = defaultProduct
	ps.nextID++

	return ps
}

func (ps *ProductStorage) GetAll() []models.Product {
	var products []models.Product
	for _, product := range ps.products {
		products = append(products, product)
	}
	return products
}

func (ps *ProductStorage) Create(product models.CreateProductRequest) models.Product {
	newProduct := models.Product{
		ID:          ps.nextID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		SKU:         product.SKU,
		Category:    product.Category,
		CreatedAt:   time.Now(),
	}
	ps.products[ps.nextID] = newProduct
	ps.nextID++
	return newProduct
}
