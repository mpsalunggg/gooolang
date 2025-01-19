package main

import "fmt"

func main(){
	// Interface any

	var secret interface{}

	secret = "mps"

	fmt.Println(secret)

	secret = []string{"apple", "mangga", "pineapple"}

	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"name": "mps",
		"grade": 2,
		"fruits": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)
}