package main

import "fmt"

func exercise1(operator string, a int, b int) int {
	switch operator {
	case "+":
		fmt.Println("Tambah:", a+b)
		return a + b
	case "-":
		fmt.Println("Kurang:", a-b)
		return a - b
	case "*":
		fmt.Println("Kali:", a*b)
		return a * b
	case "/":
		fmt.Println("Bagi:", a/b)
		return a / b
	default:
		fmt.Println("Operator tidak valid")
		return 0
	}
}

func exercise2(n int) string {
	genap := "Genap bro"
	ganjil := "Ganjil boskuu"
	if n%2 == 0 {
		fmt.Println(genap)
		return genap
	}
	fmt.Println(ganjil)
	return ganjil
}

func exercise3(nilai int) string {
	if nilai >= 80 {
		fmt.Println("A")
		return "A"
	} else if nilai >= 70 {
		fmt.Println("B")
		return "B"
	} else if nilai >= 60 {
		fmt.Println("C")
		return "C"
	} else {
		fmt.Println("D")
		return "D"
	}
}

func main() {
	// exercise 1
	exercise1("+", 1, 2)
	exercise1("-", 1, 2)
	exercise1("*", 1, 2)
	exercise1("/", 1, 2)
	exercise1("^", 1, 2)
	exercise1("%", 1, 2)
	exercise1("&", 1, 2)

	// exercise 2
	exercise2(1)
	exercise2(2)

	// exercise 3
	exercise3(80)
	exercise3(70)
	exercise3(60)
	exercise3(50)
}
