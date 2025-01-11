package main

import "fmt"


type People struct {
	Name string
	Age int
}

func (p People) sayHello() {
    fmt.Println("Hello", p.Name)
}

func main(){
	people := People{"mps", 23}
	people.sayHello()
}