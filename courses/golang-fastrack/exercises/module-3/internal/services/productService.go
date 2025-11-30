package services

import (
	"fmt"
	"math/rand"
	"module-3/internal/models"
	"strconv"
)

var products []models.Product

func generateRandomId() string {
	return strconv.Itoa(rand.Intn(1000000))
}

func CreateProduct(name string, price float64) models.Product {
	product := models.Product{
		Id:    generateRandomId(),
		Name:  name,
		Stock: 0,
		Price: price,
	}

	products = append(products, product)
	return product
}

func GetProductById(id string) *models.Product {
	for i := range products {
		if products[i].Id == id {
			return &products[i]
		}
	}
	return nil
}

func UpdateProduct(id, name string, price float64) models.Product {
	product := GetProductById(id)

	if product.Id == "" {
		return models.Product{}
	}

	product.Name = name
	product.Price = price
	return *product
}

func UpdateStokProduct(id string, stock int) *models.Product {
	product := GetProductById(id)
	if product.Id == "" {
		return nil
	}
	product.Stock = stock
	return product
}

func DeleteProduct(id string) *models.Product {
	for i := range products {
		if products[i].Id == id {
			deletedProduct := products[i]
			fmt.Println("before, ", products[:i])
			fmt.Println("after, ", products[i+1:])
			products = append(products[:i], products[i+1:]...)
			return &deletedProduct
		}
	}
	return nil
}

func GetAllProducts() []models.Product {
	return products
}
