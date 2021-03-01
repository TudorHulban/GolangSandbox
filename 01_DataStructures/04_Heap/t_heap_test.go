package main

import (
	"fmt"
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

	h := NewPriority()
	h.Insert(&g1).Insert(&g2)

	require.Len(t, h.State, 2)

	fmt.Println(h.Values())
}
