package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(1, 5)
	fmt.Println(result)
}
