package main

import "fmt"

func main() {
	fmt.Println("POINTER")
	// Basic Pointer with reference and deference
	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println("numberA:", numberA)
	fmt.Println("numberA:", &numberA)
	fmt.Println("numberB:", *numberB)
	fmt.Println("numberB:", numberB)
	fmt.Println("======")
	// Pointer changes value
	*numberB = 10
	fmt.Println("numberA:", numberA)
	fmt.Println("numberA:", &numberA)
	fmt.Println("numberB:", *numberB)
	fmt.Println("numberB:", numberB)
	fmt.Println("======")
	// Pointer with parameter
	numberC := 4
	fmt.Println("Before", numberC)
	changeValue(&numberC)
	fmt.Println("After", numberC)
	fmt.Println("numberC:", &numberC)
	fmt.Println("======")
	// Swap value
	numberD := 4
	numberE := 5
	fmt.Println("Before", numberD, numberE)
	swapValue(&numberD, &numberE)
	fmt.Println("After", numberD, numberE)
}

func changeValue(number *int) {
	*number = 100
}

func swapValue(numA, numB *int) {
	temp := *numA
	*numA = *numB
	*numB = temp
}
