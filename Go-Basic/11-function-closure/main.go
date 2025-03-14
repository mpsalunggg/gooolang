package main

import "fmt"

func main() {
	fmt.Println("FUNCTION CLOSURE")
	getMinMax()
	countNum()
	countNumSlice()
	fmt.Println("======")
	counter := getCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println("======")
	visitNum([]int{1, 2, 3, 4, 5}, func(num int) {
		fmt.Println(num)
	})
	fmt.Println("======")

	listFruits := []string{"apple", "banana", "grape", "orange", "melon"}
	checkString := func(str string) bool {
		return len(str) > 5
	}

	filterStrings(listFruits, checkString)
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

// Immediately-Invoked Function Expression
func countNum() {
	result := func(a, b int) int {
		return a + b
	}(2, 4)

	fmt.Println(result)
}

// Immediately-Invoked Function Expression with Slice
func countNumSlice() {
	allNum := []int{2, 1, 3, 4, 1}
	count := func(num []int) int {
		var result int
		for _, e := range num {
			result += e
		}
		return result
	}(allNum)
	fmt.Printf("Total: %d\n", count)
}

// Closure as Return Value
func getCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

// Closure as Argument or Callback
func visitNum(numbers []int, callback func(int)) {
	for _, e := range numbers {
		callback(e)
	}
}

// Closure with Strings using Type Alias
type StringFilterCallback func(string) bool

func filterStrings(strings []string, callback StringFilterCallback) []string {
	var result []string
	for _, str := range strings {
		if callback(str) {
			result = append(result, str)
		}
	}
	return result
}
