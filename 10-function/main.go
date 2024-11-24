package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Putra", "Satria"}
	callName("Haiii, ", names)
}

func callName(message string, names []string) string {
	nameString := strings.Join(names, " ")
	fmt.Println(message, nameString)
	return nameString
}
