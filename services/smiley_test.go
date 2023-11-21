package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSmileys(t *testing.T) {
	validFaces := []string{":)", ";D", ":~)", ";-D"}
	result := CountSmileys(validFaces)
	assert.Equal(t, 4, result)

	invalidFaces := []string{";(", ":>", ":}", ":]"}
	result = CountSmileys(invalidFaces)
	assert.Equal(t, 0, result)
}
