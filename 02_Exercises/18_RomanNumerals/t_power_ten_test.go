package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowers(t *testing.T) {
	testCases := []struct {
		description string
		val         int
		expected    int
	}{
		{"0", 0, 1},
		{"1", 1, 10},
		{"2", 2, 100},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, powersOfTen(testCase.val), testCase.description)
	}
}
