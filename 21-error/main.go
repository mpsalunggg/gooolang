package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("start...")

	var value int
	fmt.Print("enter your number: ")
	_, err := fmt.Scanf("%d", &value)
	if err != nil {
		fmt.Println("errorrr : ", err)
		return
	}

	err = ProcessNumber(value)
	if err != nil {
		fmt.Println("damnn error bruh : ", err)
		return
	}

	fmt.Println("finish...")
}

func ProcessNumber(num int) error {
	if num < 0 {
		return errors.New("number less than 0")
	}

	fmt.Println("valid input : ", num)
	return nil
}
