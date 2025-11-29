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
	// Looping with range num
	// Error cannot range over 5 , change to use array works
	var test = []int{0, 1, 2, 3, 4} 
	for i, _ := range test {
		fmt.Println(i)
	}

	var xs = "123"
	for i, v := range xs {
    	fmt.Println("Index=", i, "Value=", v)
	}

	// Break & Continue
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Num ", i)
	}

	// labeling
	OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 2 {
                fmt.Println("Keluar dari OuterLoop")
                break OuterLoop
            }
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }
    fmt.Println("Perulangan selesai")
}
