package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPairs(t *testing.T) {
	data := []int{1, 2, 1, 1, 3, 2, 4, 2, 2}

	p1, errP1 := getPairs(2, data)
	require.Nil(t, errP1)
	assert.Equal(t, 3, p1) // pairs are for 1 - 1 pair, for 2 - 2 pairs

	p2, errP2 := getPairs(3, data)
	require.Nil(t, errP2)
	assert.Equal(t, 2, p2) // pairs are for 1 - 1 pair, for 2 - 2 pairs
}
