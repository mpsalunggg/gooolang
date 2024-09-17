package main

import "fmt"

func main(){
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	const name = "Putra"
	const age = 20

	// fmt print
	fmt.Print("Hai nama saya ", name, " dan umur saya ", age, "\n")

	// fmt println
	fmt.Println("halooo banggg")

	// multi constanta
	const (
		circle	= "Lingkarang"
		isToday bool = false
		number uint8 = 33
		floatNumber = 2.2
	)
	fmt.Println(circle, isToday, number, floatNumber)

	// multi constanta with same value
	const (
		z = "Heyyy"
		x
	)	
	fmt.Println(z, x)

	// multi constanta with one line
	const satu, dua = 1, 2
	const two, five uint8 = 2, 5
	fmt.Println(satu, dua, two, five)
}