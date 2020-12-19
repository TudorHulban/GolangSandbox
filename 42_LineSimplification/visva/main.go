package main

import (
	"log"

	"github.com/intdxdt/geom"
)

func main() {
	line := []geom.Point{{3.0, 1.6}, {3.0, 2.0}, {2.4, 2.8}, {0.5, 3.0}, {1.2, 3.2}, {1.4, 2.6}, {2.0, 3.5}, {5, 6}}
	log.Println("Line Points: ", len(line))

	visva := NewVisva(line)
	res := 0.85
	simplx := visva.Simplify(res)

	l1 := geom.NewLineString(geom.Coordinates(line))
	log.Println("Line Points l1: ", len(l1.Coordinates.Pnts))
	log.Println(l1.WKT())

	l2 := geom.NewLineString(geom.Coordinates(simplx))
	log.Println("Line Points l2: ", len(l2.Coordinates.Pnts))
	log.Println(l2.WKT())
}
