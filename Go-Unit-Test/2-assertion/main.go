package main

import "fmt"

func SayHelloPeople(name string) string {
	return "Hello, " + name
}

func GreetPeople(names []string, greeting string) []string {
	var greetings []string
	for _, name := range names {
		greetings = append(greetings, greeting+", "+name)
	}
	return greetings
}

func main() {
	fmt.Println(SayHelloPeople("John"))
	fmt.Println(GreetPeople([]string{"John", "Jane", "Jim"}, "Hello"))
}
