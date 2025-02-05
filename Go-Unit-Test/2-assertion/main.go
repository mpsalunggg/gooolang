package main

import "fmt"

func SayHelloPeople(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(SayHelloPeople("John"))
}
