package main

import (
	"fmt"
	"strings"
)

func main() {
	// Interface any

	var secret interface{}

	secret = "mps"

	fmt.Println(secret)

	secret = []string{"apple", "mangga", "pineapple"}

	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	data := map[string]interface{}{
		"name":   "mps",
		"grade":  2,
		"fruits": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)

	data2 := map[string]any{}

	fmt.Println(data2)
	data2 = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data2)
	data2["say"] = "hallooo"
	fmt.Println(data2["say"])

	fmt.Println("=======================")
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

}
