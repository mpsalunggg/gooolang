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

	// Anonymous Struct
	var personB = struct {
		Name string
		Age  int
	}{
		Name: "Jane",
		Age:  22,
	}
	fmt.Println(personB)

	// All Initialize Struct
	var s1 = Person{"Doe", 24}
	fmt.Println(s1)
	var s2 = Person{Age: 24, Name: "Doe"}
	fmt.Println(s2)
	var s3 = Person{"Doe", 24}
	fmt.Println(s3)

	fmt.Println("=====================================")
	// Struct with pointer
	var personC = &Person{"Doe", 24}
	fmt.Println(personC)
	fmt.Println(personC.Name)
	fmt.Println(personC.Age)

	fmt.Println("=====================================")
	var person1 = Person{Name: "Doe", Age: 24}

	var person2 *Person = &person1
	fmt.Println("student 1, name :", person1.Name)
	fmt.Println("student 4, name :", person2.Name)

	person2.Name = "Smith"
	fmt.Println("student 1, name :", person1.Name)
	fmt.Println("student 4, name :", person2.Name)

}
