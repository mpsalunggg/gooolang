package main

import (
	"bufio"
	"fmt"
	"learn-golang/21-error/example2/errors"
	"os"
	"strconv"
	"strings"
)

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

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	people.Name = strings.TrimSpace(name)

	fmt.Print("Enter your balance: ")
	balanceInput, _ := reader.ReadString('\n')
	balanceInput = strings.TrimSpace(balanceInput)
	balance, err := strconv.ParseFloat(balanceInput, 64)
	if err != nil {
		fmt.Println(errors.ErrorInputBalance().Error())
		return
	}
	people.Balance = balance

	fmt.Printf("welcome to the %s restaurant %s :)\n", restaurant.Name, people.Name)
}
