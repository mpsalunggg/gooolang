package example3

import "fmt"

type BankAccount interface {
	Deposit(amount float64)
	WithDraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	Balance float64
	Owner   string
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	fmt.Printf("Deposit %.2f berhasil. Saldo saat ini: %.2f\n", amount, a.Balance)
}

func (a *Account) WithDraw(amount float64) error {
	if amount > a.Balance {
		return fmt.Errorf("saldo tidak cukup: saldo saat ini %.2f", a.Balance)
	}
	a.Balance -= amount
	fmt.Printf("Penarikan %.2f berhasil. Saldo saat ini: %.2f\n", amount, a.Balance)
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func PrintBalance(account BankAccount) {
	fmt.Printf("Saldo saat ini: %.2f\n", account.GetBalance())
}

func Example() {
	myAccount := &Account{Owner: "Putra", Balance: 10000.0}
	// fmt.Println(*&myAccount.Owner)
	fmt.Println(myAccount.Owner)

	PrintBalance(myAccount)

	myAccount.WithDraw(2000.0)

	myAccount.Deposit(5000.0)

	err := myAccount.WithDraw(200000.00)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
