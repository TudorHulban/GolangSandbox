package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	m := identifyString("saa")
	s := orderMapString(m)

	require.Len(t, s, 2)
	require.Equal(t, s[0].Value, 1)
	require.Equal(t, s[1].Value, 2)
}
