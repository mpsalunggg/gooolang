package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHelloPeople(t *testing.T) {
	exampleName := "Yuhu"
	result := SayHelloPeople(exampleName)

	assert.Equal(t, "Hello, "+exampleName, result, "Result should be Hello, "+exampleName)
	assert.Contains(t, result, "Hello, "+exampleName, "Result should contain Hello, "+exampleName)

	// if the assertion is not passed, the test will be stopped and the message will be printed
	fmt.Println("TestSayHelloPeople with assertion passed")
}

func TestSayHelloPeopleWithRequire(t *testing.T) {
	exampleName := "Yuhu"
	result := SayHelloPeople(exampleName)

	require.Equal(t, "Hello, "+exampleName, result, "Result should be Hello, "+exampleName)

	// if the require is not passed, the test will be stopped and the message will be printed
	fmt.Println("TestSayHelloPeopleWithRequire passed")
}

func TestGreetPeople(t *testing.T) {
	names := []string{"John", "Jane", "Jim"}
	greeting := "Hello"
	result := GreetPeople(names, greeting)

	assert.Equal(t, []string{"Hello, John", "Hello, Jane", "Hello, Jim"}, result)

	assert.Contains(t, result, "Hello, John")
	assert.Contains(t, result, "Hello, Jane")
	assert.Contains(t, result, "Hello, Jim")
}

func TestSkip(t *testing.T) {
	skip := true
	if skip {
		t.Skip("Skip this test")
	}

	assert.Equal(t, "Hello, John", SayHelloPeople("John"))
}

func TestMain(m *testing.M) {
	fmt.Println("Befor testing")

	m.Run()

	fmt.Println("After testing")
}

func TestSum(t *testing.T) {
	t.Run("2+3", func(t *testing.T) {
		if result := Sum(2, 3); result != 5 {
			t.Errorf("expected 5, got %d", result)
		}
	})

	t.Run("4+6", func(t *testing.T) {
		if result := Sum(4, 6); result != 10 {
			t.Errorf("expected 10, got %d", result)
		}
	})
}