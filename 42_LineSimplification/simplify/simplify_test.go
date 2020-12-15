package simplify

import (
	"log"
	"testing"
)

var pts = [][]float64{{1, 2}, {2, 3}, {3, 1}, {4, 10}, {5, 25}}

func TestZeroTol(t *testing.T) {
	simplex, err := Floats(pts, 0, true)
	if err != nil {
		t.Error(err)
	}
	if len(pts) != len(simplex) {
		t.Error("tolerance not working", len(pts), len(simplex))
	}
	log.Println(simplex)
}

func TestTwoTol(t *testing.T) {
	simplex, err := Floats(pts, 2, true)
	if err != nil {
		t.Error(err)
	}
	if len(pts) == len(simplex) {
		t.Error("tolerance not working", len(pts), len(simplex))
	}
	log.Println(simplex)
}

func TestTenTol(t *testing.T) {
	simplex, err := Floats(pts, 10, true)
	if err != nil {
		t.Error(err)
	}
	if len(simplex) != 2 {
		t.Error("tolerance not working", len(pts), len(simplex))
	}
	log.Println(simplex)
}
