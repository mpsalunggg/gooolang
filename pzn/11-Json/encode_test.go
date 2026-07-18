package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	logger := []string{"Google", "Gojek", "Traveloka"}

	bytes, err := json.Marshal(logger)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(bytes))
}
