package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)

	if result != 4 {
		t.Error("Number should be same the result")
	}

	// this will be executed when the test is passed
	fmt.Println("Test Add passed")
}

func TestAdd2(t *testing.T) {
	result := Add(1, 2)

	if result != 4 {
		t.Fatal("Number should be same the result")
	}

	// this will be executed when the test is failed
	fmt.Println("Test Add2 passed")
}

func TestAdd3(t *testing.T) {
	result := Add(1, 2)

	if result != 4 {
		t.Fail()
	}

	// this will be executed when the test is failed
	fmt.Println("Test Add3 passed")
}

func TestAdd4(t *testing.T) {
	result := Add(1, 2)

	if result != 4 {
		t.FailNow()
	}

	// this not will be executed when the test is failed
	fmt.Println("Test Add4 passed")
}
