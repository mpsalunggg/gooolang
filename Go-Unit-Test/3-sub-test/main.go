package main

import "fmt"


func SumTwoNumber(num1, num2 int) int {
  return num1 + num2
}

func main(){
	fmt.Println(SumTwoNumber(2,3))
	fmt.Println(SumTwoNumber(3,5))
}