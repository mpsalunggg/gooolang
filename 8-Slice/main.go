package main

import "fmt"

func main() {
	// slice same with array, but slice dont have spesific length
	var fruits = []string{"bananan", "grape", "etc"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)


	// Relation Array with Slice
	var new_fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(new_fruits[0:2])
	fmt.Println(new_fruits[4:4])
	// fmt.Println(new_fruits[4:0]) error because 0 is < 4
	fmt.Println(new_fruits[:]) //all data on array
	fmt.Println(new_fruits[:2])
	fmt.Println(new_fruits[2:])

}
