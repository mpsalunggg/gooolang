package main

import "fmt"

func main() {
	// slice same with array, but slice dont have spesific length
	var fruits = []string{"bananan", "grape", "etc"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)

}
