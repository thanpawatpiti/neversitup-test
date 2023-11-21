package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchChar(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"abc", []string{"abc", "acb", "bac", "bca", "cba", "cab"}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := []string{}
			SwitchChar(tc.input, 0, &result)

			assert.Equal(t, tc.expected, result)
		})
	}
}
