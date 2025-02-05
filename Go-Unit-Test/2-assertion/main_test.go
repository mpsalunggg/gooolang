package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHelloPeople(t *testing.T) {
	exampleName := "Yuhu"
	result := SayHelloPeople(exampleName)

	assert.Equal(t, "Hello, "+exampleName, result, "Result should be Hello, "+exampleName)
	assert.Contains(t, result, "Hello, "+exampleName, "Result should contain Hello, "+exampleName)
	// assert.NotContains(t, result, "Hello, "+exampleName, "Result should not contain Hello, "+exampleName)
}
