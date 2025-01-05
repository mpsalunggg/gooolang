package main

import "fmt"

func main() {
	fmt.Println("STRUCT")
	// Basic Struct
	type Person struct {
		Name string
		Age  int
	}

	var personA Person
	personA.Name = "John"
	personA.Age = 20
	fmt.Println(personA)
}