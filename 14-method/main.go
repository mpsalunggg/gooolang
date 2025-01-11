package main

import (
	"fmt"
	"strings"
)


type People struct {
	Name string
	Age int
}

func (p People) sayHello() {
    fmt.Println("Hello", p.Name)
}

func (p People) getFirstName(i int) string {
	return strings.Split(p.Name, " ")[i - 1]
}

func main(){
	people := People{"putra satria", 23}
	people.sayHello()

	firstName := people.getFirstName(1)
	fmt.Println("First Name:", firstName)
}