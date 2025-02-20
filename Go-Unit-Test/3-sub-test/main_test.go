package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumTwoNumber(t *testing.T) {
	t.Run("2 + 3", func(t *testing.T) {
		result := SumTwoNumber(2, 3)
		require.Equal(t, 5, result, "The sum of 2 and 3 is not equal to 5")
	})

	t.Run("5 + 2", func(t *testing.T) {
		result := SumTwoNumber(5, 2)
		require.Equal(t, 7, result, "The sum of 2 and 3 is not equal to 5")
	})
}
