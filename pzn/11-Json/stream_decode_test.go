package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStreamDecoderFile(t *testing.T) {
	file, err := os.Open("Customer.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	customer := &Customer{}
	err = decoder.Decode(customer)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestStreamDecoderMultiple(t *testing.T) {
	jsonString := `{"first_name":"Budi","last_name":"Nugroho","age":25}
{"first_name":"Sari","last_name":"Wijaya","age":30}
{"first_name":"Andi","last_name":"Saputra","age":22}`

	reader := strings.NewReader(jsonString)
	decoder := json.NewDecoder(reader)

	for decoder.More() {
		customer := &Customer{}

		err := decoder.Decode(customer)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(customer.FirstName, customer.LastName, customer.Age)
	}
}
