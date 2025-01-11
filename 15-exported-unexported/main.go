package main

import (
	"fmt"
	"learn-golang/15-exported-unexported/people"
)

func main() {
	// Example of exported function / public
	sayHello := people.SayHello("Putra")
	fmt.Println(sayHello)
	fmt.Println()
	// Example of unexported function / private
	// fmt.Println(people.sayHello("Putra"))
}