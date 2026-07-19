package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type People struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJSONTagEncode(t *testing.T) {
	people := People{
		Id:   "1",
		Name: "mps",
		Age:  20,
	}

	bytes, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	stringJson := `{"id":"1","name":"mps","age":20}`
	jsonBytes := []byte(stringJson)

	people := &People{}

	json.Unmarshal(jsonBytes, people)

	fmt.Println(people)
}
