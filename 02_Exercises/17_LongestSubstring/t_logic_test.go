package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasRepeat(t *testing.T) {
	require.True(t, noRepeating("abc"))
	require.False(t, noRepeating("xxx"))
}

func TestGetLongest(t *testing.T) {
	testCases := []struct {
		description string
		val         string
		expected    string
	}{
		{"repeating", "aa", "a"},
		{"simple", "aab", "ab"},
		{"general", "abcabcbb", "abc"},
		{"substring", "pwwkew", "wke"},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, walkString(testCase.val), testCase.description)
	}
}
