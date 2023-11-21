package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOdd(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 2, 1, 3, 4, 4}, 3},
		{[]int{5, 7, 5, 7, 2, 2, 1}, 1},
		{[]int{9, 9, 9, 9, 5, 5, 7}, 7},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := FindOdd(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
