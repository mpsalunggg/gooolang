package main

import (
	"fmt"
	"module-3/internal/services"
)

func main() {
	fmt.Println("Hello World")
	product := services.CreateProduct("Product 1", 100000)
	product2 := services.CreateProduct("Product 2", 4000)
	services.CreateProduct("Product 3", 5000)
	fmt.Println(product)
	fmt.Println(product2)

	// GET Product by id example
	productById := services.GetProductById(product.Id)
	fmt.Println("Product by id :\n", *productById)

	// UPDATA Product
	updatedProduct := services.UpdateProduct(product.Id, "Product 1 Updated", 2000)
	fmt.Println("Update product :\n", updatedProduct)

	// UPDATE Stok product
	updatedStokProduct := services.UpdateStokProduct(product.Id, 10)
	fmt.Println("Update stock product :\n", updatedStokProduct)

	// DELETE Product
	deletedProduct := services.DeleteProduct(product2.Id)
	fmt.Println("Delete product :\n", *deletedProduct)

	// all products
	allProducts := services.GetAllProducts()
	fmt.Println("All products")
	for _, product := range allProducts {
		fmt.Println("Product ID :", product.Id, "Product Name :", product.Name, "Product Stock :", product.Stock, "Product Price :", product.Price)
	}
}
