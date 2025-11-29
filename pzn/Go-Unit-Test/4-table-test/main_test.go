package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolindromeTableTest(t *testing.T) {
	tests := []struct {
		text     string
		request  string
		expected bool
	}{
		{
			text:     "makam",
			request:  "makam",
			expected: true,
		},
		{
			text:     "katak",
			request:  "katak",
			expected: true,
		},
		{
			text:     "mau",
			request:  "mau",
			expected: false,
		},
	}

	for _, val := range tests {
		t.Run(val.text, func(t *testing.T) {
			result := IsPalindrome(val.request)

			require.Equal(t, val.expected, result)
		})
	}
}
