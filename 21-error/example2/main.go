package main

import "fmt"

type People struct {
	Name    string
	Balance float64
}

type Menu struct {
	Name  string
	Price float64
}

type Restaurant struct {
	Name  string
	Menus []Menu
}

func main() {
	people := People{"Putra", 100000}
	restaurant := Restaurant{
		Name: "Nagaya",
		Menus: []Menu{
			{Name: "Nasi Goreng", Price: 20000},
			{Name: "Mie Ayam", Price: 10000},
			{Name: "Ikan Bakar", Price: 25000},
		},
	}

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &people.Name)
	fmt.Print("Enter your balance: ")
	fmt.Scanf("%f", &people.Balance)

	fmt.Printf("welcome to the %s restaurant %s :)\n", restaurant.Name, people.Name)
}
