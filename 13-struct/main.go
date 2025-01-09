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

	// Embeded Struct
	type Address struct {
		City, Province, Country string
	}
	type Contact struct {
		Phone, Email string
		Address
	}

	var contact1 = Contact{}
	contact1.Phone = "0821-2222-3333"
	contact1.Email = "people@gmail.com"
	contact1.Address.City = "Jakarta"
	contact1.Address.Province = "DKI Jakarta"
	contact1.Address.Country = "Indonesia"

	fmt.Println(contact1)
	fmt.Println(contact1.Phone)

	// Filling value sub struct
	var contact2 = Contact{
		Phone: "0821-2222-3333",
		Email: "people@gmail.com",
		Address: Address{
			City:     "Jakarta",
			Province: "DKI Jakarta",
			Country:  "",
		},
	}

	fmt.Println(contact2)

	fmt.Println("=====================================")
	type People struct {
		Name string
		Age  int
	}

	type Student struct {
		People
		Grade int
	}

	var people = People{"Doe", 24}
	var teacher = Student{People: people, Grade: 100}
	fmt.Println(teacher)

	// Anonymous struct
	var teacher2 = struct {
		People
		Grade int
	}{}

	teacher2.People = People{"Smith", 24}
	teacher2.Grade = 100

	teacher3 := struct {
		People
		Grade int
	}{
		People: People{"Mps", 24},
		Grade:  20,
	}
	fmt.Println(teacher2)
	fmt.Println(teacher3)

	fmt.Println("=====================================")
	// struct with
	type Fruits struct {
		Name  string
		Color string
	}

	allFruits := []Fruits{
		{"Apple", "Red"},
		{"Banana", "Yellow"},
		{"Grape", "Purple"},
	}

	for index, fruit := range allFruits {
		fmt.Println(index+1, ".", fruit.Name, ":", fruit.Color)
	}

	fmt.Println(allFruits)
}
