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

func SkipTest(t *testing.T) {
	skip := true
	if skip {
		t.Skip("Skip this test")
	}

	assert.Equal(t, "Hello, John", SayHelloPeople("John"))
}
