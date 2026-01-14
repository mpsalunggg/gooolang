package main

import "fmt"

func main(){
	// Basic Variable
	var firstName, lastName string = "putra", "satria";  
	fmt.Println(firstName + lastName)

	test := "test"
	test2 := "test2"

	fmt.Printf("halo %s %s!\n", test, test2)

	// Underscore Variable
	_ = "testt"
	name , _ := "putra", "satria"
	fmt.Println(name)

	// Keyword new -> pointer -> * = tanda asterisk
	fullname := new(string)
	fmt.Println(fullname)
	fmt.Println(*fullname) // string kosong

	*fullname = "haii"
	fmt.Println(*fullname) // string haii

}