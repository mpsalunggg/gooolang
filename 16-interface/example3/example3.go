package example3

import "fmt"

type BankAccount interface {
	Deposit(amount float64)
	WithDraw(amount float64) error
}

type Account struct {
	Balance float64
	Owner   string
}

func Example() {
	myAccount := &Account{Balance: 10000, Owner: "Putra"}
	// fmt.Println(*&myAccount.Owner)
	fmt.Println(myAccount.Owner)
}
