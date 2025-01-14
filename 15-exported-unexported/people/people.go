package people

import (
	f "fmt" 
)

func SayHello(name string) string {
	return "Hello " + name
}

func sayHello(name string) string {
	return "Hello " + name
}

func Whatsup(str string) {
	f.Println(str)
}