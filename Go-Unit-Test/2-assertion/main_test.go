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
