package main

import "fmt"

func main() {
	fmt.Println("MAP")
	// var people map[string]int
	people := map[string]string{
		"name":    "putra",
		"stambuk": "f55120009",
	}

	people["age"] = "2"
	people["age"] = "5"
	fmt.Println(people)

	// month := make(map[int]string)
	// month2 := *new(map[string]int)

	month := map[int]string{
		1: "January",
		2: "February",
	}

	fmt.Println(month)
	// fmt.Println(month2)

	var chicken = map[string]int{
		"katsu":  2,
		"rembo":  3,
		"geprek": 5,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	// Delete Object Key
	var cow = map[string]string{}
	cow["1"] = "Test"
	cow["2"] = "Test2"

	fmt.Println(cow)
	delete(cow, "1")

	// Detect Item With Key
	value, isExist := chicken["katsu"]

	if isExist {
		fmt.Println("value:", value)
	} else {
		fmt.Println("Item is not exist")
	}
}
