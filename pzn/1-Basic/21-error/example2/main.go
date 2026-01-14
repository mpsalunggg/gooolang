package main

import (
	"bufio"
	"fmt"
	"learn-golang/21-error/example2/errors"
	"os"
	"strconv"
	"strings"

	"github.com/leekchan/accounting"
)

type People struct {
	Name    string
	Balance float64
}

type Menu struct {
	ID    int
	Name  string
	Price float64
}

type RestaurantService interface {
	FindByID(id int) (*Menu, error)
	Purchase(people *People, menu *Menu) error
}

type Restaurant struct {
	Name  string
	Menus []Menu
}

func (r *Restaurant) FindByID(id int) (*Menu, error) {
	for _, menu := range r.Menus {
		if menu.ID == id {
			return &menu, nil
		}
	}
	return nil, errors.ErrorInvalidMenu()
}

func (r *Restaurant) Purchase(people *People, menu *Menu) error {
	if people.Balance < menu.Price {
		return errors.ErrorBalanceNotEnough()
	}

	people.Balance -= menu.Price
	fmt.Println("Thank you for purchasing:)")
	return nil
}

func main() {
	people := People{}
	restaurant := Restaurant{
		Name: "Nagaya",
		Menus: []Menu{
			{ID: 1, Name: "Nasi Goreng", Price: 20000},
			{ID: 2, Name: "Mie Ayam", Price: 10000},
			{ID: 3, Name: "Ikan Bakar", Price: 25000},
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

	fmt.Println("List menu : ")
	for i, val := range restaurant.Menus {
		fmt.Printf("%d. %s \t: %f\n", i+1, val.Name, val.Price)
	}

	fmt.Printf("welcome to the %s restaurant %s :)\n", restaurant.Name, people.Name)

	fmt.Print("\nEnter the menu ID you want to buy: ")
	orderInput, _ := reader.ReadString('\n')
	orderInput = strings.TrimSpace(orderInput)
	menuID, err := strconv.Atoi(orderInput)
	if err != nil {
		fmt.Println(errors.ErrorInvalidMenu().Error())
		return
	}

	menu, err := restaurant.FindByID(menuID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = restaurant.Purchase(&people, menu)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("==============================")

	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 0,
		Thousand:  ".",
		Decimal:   ",",
	}
	fmt.Printf("your money left %s\n", ac.FormatMoney(people.Balance))
}
