package main

import "fmt"

func main() {
	fmt.Println("FUNCTION CLOSURE")
	getMinMax()
}

// Closure with variable
func getMinMax() {
	resultMinMax := func(num []int) (int, int) {
		var min, max int
		for i, e := range num {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	numbers := []int{3, 12, 4, 4, 1, 3}
	min, max := resultMinMax(numbers)
	fmt.Println("data min: ", min)
	fmt.Println("data max: ", max)
}
