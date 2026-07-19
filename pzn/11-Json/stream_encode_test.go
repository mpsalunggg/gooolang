package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoderFile(t *testing.T) {
	customer := Customer{
		FirstName:  "Budi",
		MiddleName: "Adi",
		LastName:   "Nugroho",
		Age:        25,
		Married:    false,
		Hobbies:    []string{"Coding", "Gaming"},
	}

	file, err := os.Create("CustomerOut.json")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(customer)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Berhasil encode ke file CustomerOut.json")
}

func TestEncoderIndent(t *testing.T){
	customer := Customer{
		FirstName: "Mps",
		MiddleName: "Tes",
		LastName: "Al",
		Age: 20,
		Married: false,
		Hobbies: []string{"Basketball", "MiniSoccer"},
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(customer)
	if err != nil {
		t.Fatal(err)
	}
}