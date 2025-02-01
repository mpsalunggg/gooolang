package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Number should be same the result")
	}
}
