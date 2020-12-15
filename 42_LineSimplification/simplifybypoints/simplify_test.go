package simplify

import (
	"log"
	"testing"
)

var points = []Point{{X: 1, Y: 2}, {X: 2, Y: 3}, {X: 3, Y: 1}, {X: 4, Y: 10}, {X: 5, Y: 25}, {X: 6, Y: 11}}

func TestZeroTol2(t *testing.T) {
	simplex, err := Points(points, 0, true)
	if err != nil {
		t.Error(err)
	}
	if len(points) != len(simplex) {
		t.Error("tolerance not working", len(points), len(simplex))
	}
	log.Println(simplex)
}

func Test__TwoTol12(t *testing.T) {
	simplex, err := Points(points, 2, true)
	if err != nil {
		t.Error(err)
	}
	if len(points) == len(simplex) {
		t.Error("tolerance not working", len(points), len(simplex))
	}
	log.Println(simplex)
}

func TestTenTol2(t *testing.T) {
	simplex, err := Points(points, 10, true)
	if err != nil {
		t.Error(err)
	}
	if len(simplex) != 3 {
		t.Error("tolerance not working", len(points), len(simplex))
	}
	log.Println(simplex)
}
