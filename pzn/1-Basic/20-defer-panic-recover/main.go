package main

import "fmt"

func Finish() {
	fmt.Println("finish...")
}

func CheckModule(num int) {
	if num%2 != 0 {
		panic("number is not modulo")
	} else {
		fmt.Println("number is modulo")
	}
}

func main() {
	defer Finish()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("something error : ", r)
		}
	}()
	fmt.Println("startt...")

	CheckModule(3)
}
