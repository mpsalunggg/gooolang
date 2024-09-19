package main

import (
	"fmt"
	"strings"
)

func main() {
	// Condition with temporary
	var age = 20
	if newAge := age + 2; newAge > 20 {
		fmt.Println("Old")
	} else {
		fmt.Println("Young")
	}

	// Switch case with nested if else
	name := "putra satria"

	switch s := strings.Contains(name, "a"); s {
	case true:
		fmt.Println("letter 'a' in the name")
		count := strings.Count(name, "a")
		fmt.Println("count letter a", count)
		if count > 1 {
			fmt.Println("> 1")
		} else {
			fmt.Println("< 1")
		}
	case false:
		fmt.Println("letter 'a is not in the name")
	}
}