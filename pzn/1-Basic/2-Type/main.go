package main

import "fmt"

func main(){
	// type data
	var positifeNumber uint8 = 89
	var negativeNumber = -1243423644
	var desimalNumber = 2.2
	var exist bool = true
	var message string = "Halooo"

	// d for decimal number
	fmt.Printf("Number: %d\n", positifeNumber)
	fmt.Printf("Number: %d\n", negativeNumber)
	// f for float number
	fmt.Printf("Number: %f\n", desimalNumber)
	fmt.Printf("Number: %.3f\n", desimalNumber)
	// t for boolean
	fmt.Printf("Boolean: %t \n", exist)
	// s for string
	fmt.Printf("String: %s \n", message)
	
}