package main

import "fmt"

func main() {
	// Normal looping
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Looping with argumen
	var num = 0
	for num < 3 {
		fmt.Println("Num", num)
		num++
	}

	// Looping without argument and condition
	var a = 0

	for {
		fmt.Println("Angka", a)

		a++
		if a == 5 {
			break
		}
	}
}
