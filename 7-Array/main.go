package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	// names[4] = "ez" // out off bound

	// first initialization
	var animal = [2]string{"monkey", "cow"}

	fmt.Println("Jumlah element \t\t", len(animal))
	fmt.Println("Isi semua element \t", animal)

	// without the number of elements
	var num = [...]int{2, 3, 4, 5}
	fmt.Println("jumlah elemen \t:", len(num))
	fmt.Println(num[0])

	// Array multindimention
	// var a = [2][3]int{[3]int{2, 3, 4}, [3]int{2, 3, 4}} redundant
	var a = [2][3]int{{2, 3, 4}, {2, 3, 4}}
	fmt.Println(a[0][1])

	// Array looping
	var fruits = [3]string{"Banana", "Apple", "Grape"}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
}
