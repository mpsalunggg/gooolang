package main

import "fmt"

func main() {
	// Operator Arithmetic
	const value = (((2+2)%2)*2 - 2) / 2
	fmt.Println(value)

	// Operator Comparison
	// == -> equal
	// != -> not equal
	// etc same like js

	var num1 = 1
	var num2 = 1

	var andOperator = num1+num1 == 2 && num1+num2 == 1
	var orOperator = num1-num2 == 2 || num1+num2 == 2
	fmt.Printf("(%t) \n", andOperator)
	fmt.Printf("(%t) \n", orOperator)
	fmt.Printf("(%t) \n", !orOperator)
}