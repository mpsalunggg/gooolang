package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street  string `json:"street"`
	Country string `json:"country"`
}

func TestJsonArrayEncode(t *testing.T) {
	addresses := []Address{
		{Street: "Jl. Merdeka", Country: "Indonesia"},
		{Street: "Jl. Sudirman", Country: "Indonesia"},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `[{"street":"Jl. Merdeka","country":"Indonesia"},{"street":"Jl. Sudirman","country":"Indonesia"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(addresses)
	for _, address := range *addresses {
		fmt.Println(address.Street, "-", address.Country)
	}
}
