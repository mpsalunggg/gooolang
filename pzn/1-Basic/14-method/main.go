package main

import (
	"fmt"
	"strings"
)


type People struct {
	Name string
	Age int
}

type Animal struct {
	Name string
	Type string
}

func (p People) sayHello() {
    fmt.Println("Hello", p.Name)
}

func (p People) getFirstName(i int) string {
	return strings.Split(p.Name, " ")[i - 1]
}

func (a Animal) changeName(n string){
	a.Name = n
	fmt.Println("Not change -> " , a.Name)
}

func (a *Animal) changeNamePointer(n string){
	a.Name = n
	fmt.Println("Change -> " , a.Name)
}

func main(){
	people := People{"putra satria", 23}
	people.sayHello()

	firstName := people.getFirstName(1)
	fmt.Println("First Name:", firstName)

	fmt.Println("=====================================")

	animal := Animal{"Dog", "Mammal"}
	animal.changeName("Cat")
	fmt.Println(animal.Name)
	animal.changeNamePointer("Rabbit")
	fmt.Println(animal.Name)

}