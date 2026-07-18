package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Testt",
		MiddleName: "Putra",
		LastName:   "Satria",
		Age:        25,
		Married:    false,
		Hobbies:    []string{"Coding", "Gaming"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(bytes))
}
