package main

import "fmt"

func main() {
	// slice same with array, but slice dont have spesific length
	var fruits = []string{"bananan", "grape", "etc"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)

	// Relation Array with Slice
	var new_fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(new_fruits[0:2])
	fmt.Println(new_fruits[4:4])
	// fmt.Println(new_fruits[4:0]) error because 0 is < 4
	fmt.Println(new_fruits[:]) //all data on array
	fmt.Println(new_fruits[:2])
	fmt.Println(new_fruits[2:])

	// Slice references
	var animals = []string{"monkey", "lion", "cat", "turtle"}
	var animals1 = animals[0:3] // monkey, lion, cat
	var animals2 = animals[1:4] // lion, cat, turtle

	var animals12 = animals1[0:2] // monkey, lion
	var animals22 = animals2[0:1] // lion

	fmt.Println(animals1)
	fmt.Println(animals2)
	fmt.Println(animals12)
	fmt.Println(animals22)

	// len & cap
	fmt.Println(len(animals)) //4
	fmt.Println(cap(animals)) //4

	fmt.Println(len(animals1)) //3
	fmt.Println(cap(animals1)) //4

	fmt.Println(len(animals12)) //2
	fmt.Println(cap(animals12)) //4

	// append slice
	var people = []string{"Andi", "Aba", "Aco"}
	var new_people = append(people, "Ali")
	fmt.Println(new_people)

	// another example append
	var object = []string{"Besi", "Garpu", "Tali"}
	var object_new = object[0:2]

	fmt.Println("append", cap(object_new))
	fmt.Println("append", len(object_new))

	var object_new_2 = append(object_new, "Benang")

	fmt.Println(object)       // Besi Garpu Benang
	fmt.Println(object_new)   // Besi Garpu
	fmt.Println(object_new_2) // Besi Garpu Benang

	dst := make([]string, 3)
	src := []string{"Ali", "Banee", "Mps"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
