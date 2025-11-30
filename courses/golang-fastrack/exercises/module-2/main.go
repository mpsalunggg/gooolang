package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

func findUserById(id int, users []User) (User, bool) {
	for _, user := range users {
		if user.Id == id {
			return user, true
		}
	}
	return User{}, false
}

func exercise1() {
	user1 := []User{
		{Id: 1, Name: "Ahmad", Email: "ahmad@mail.com"},
		{Id: 2, Name: "Budi", Email: "budi@mail.com"},
		{Id: 3, Name: "Charlie", Email: "charlie@mail.com"},
	}

	user, error := findUserById(2, user1)
	if error {
		fmt.Println("User found:", user)
	} else {
		fmt.Println("User not found")
	}
}

type Product struct {
	Id    string
	Name  string
	Stock int
}

type ProductMap map[string]*Product

func checkAvailableProduct(products ProductMap, productId string) bool {
	_, error := products[productId]
	return error
}

func addStock(products ProductMap, stock int, id string) {
	val := checkAvailableProduct(products, id)
	if val {
		fmt.Println("\nProduct available")
		products[id].Stock += stock
	} else {
		fmt.Println("Product not available")
	}
	fmt.Print("Now total stock on product ", id, " is ", products[id].Stock)
}

func sellProduct(products ProductMap, quantity int, id string) {
	val := checkAvailableProduct(products, id)
	if val {
		fmt.Println("\nProduct available")
		products[id].Stock -= quantity
	} else {
		fmt.Println("Product not available")
	}
	fmt.Println("Now total stock on product ", id, " is ", products[id].Stock)
}

func showAllProducts(products ProductMap) {
	for id, product := range products {
		fmt.Println("Product ID: ", id, "Product Name: ", product.Name, "Product Stock: ", product.Stock)
	}
}

func exercise2() {
	products := make(map[string]*Product)
	products["A1"] = &Product{Id: "A1", Name: "Product 1", Stock: 100}
	products["A2"] = &Product{Id: "A2", Name: "Product 2", Stock: 200}
	products["A3"] = &Product{Id: "A3", Name: "Product 3", Stock: 300}

	addStock(products, 100, "A1")
	// sell product
	sellProduct(products, 1, "A1")

	// show all products
	showAllProducts(products)
}

func main() {
	exercise1()

	exercise2()
}
