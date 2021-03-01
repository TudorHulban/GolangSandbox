package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeap(t *testing.T) {
	g1 := Generic{
		Priority: 1,
	}

	g2 := Generic{
		Priority: 7,
	}

	g3 := Generic{
		Priority: 3,
	}

	h := NewPriority()
	h.Insert(&g1).Insert(&g2).Insert(&g3)

	require.Len(t, h.State, 3)
}

func TestExtractNil(t *testing.T) {
	h := NewPriority()

	require.Nil(t, h.Extract())
}

func TestExtractAll(t *testing.T) {
	g1 := Generic{
		Priority: 1,
	}

	h := NewPriority()
	h.Insert(&g1).Extract()

	require.Len(t, h.State, 0)
}

func TestExtract(t *testing.T) {
	g1 := Generic{
		Priority: 1,
	}

	g2 := Generic{
		Priority: 7,
	}

	g3 := Generic{
		Priority: 3,
	}

	g4 := Generic{
		Priority: 5,
	}

	h := NewPriority()
	h.Insert(&g1).Insert(&g2).Insert(&g3).Insert(&g4)

	require.Equal(t, h.Extract().Priority, g2.Priority)
	require.Len(t, h.State, 3)

	require.Equal(t, h.Extract().Priority, g4.Priority)
	require.Len(t, h.State, 2)
}
