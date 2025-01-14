package main

import (
	"fmt"
	"learn-golang/15-exported-unexported/animal"
	"learn-golang/15-exported-unexported/people"
	. "learn-golang/15-exported-unexported/fruits"
)

func main() {
	// Example of exported function / public
	sayHello := people.SayHello("Putra")
	fmt.Println(sayHello)
	fmt.Println()
	// Example of unexported function / private
	// fmt.Println(people.sayHello("Putra"))

	fmt.Println("=====================================")
	// Example of exported struct / public
	animal := animal.Animal{Name: "Dog", Color: "Black"}
	fmt.Println(animal)
	// Example of unexported struct / private
	// animal2 := animal.animal{Name: "Cat", color: "White"}

	fmt.Println("=====================================")
	// Example with dot exported
	fruits := Fruits{
		Name: "Banana",
		Color: "Yellow",
		Type: "Test",
	}

	fmt.Println(fruits.Name)
	fmt.Println(fruits.Color)
	fmt.Println(fruits.Type)
	
	// Example aliases import fmt
	people.Whatsup("Hellooo gess")
}