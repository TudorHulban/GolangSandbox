package main

import (
	"log"
	"testing"

	"github.com/intdxdt/geom"
)

func TestGetSeriesFromFile(t *testing.T) {
	series, err := GetSeriesFromFile("../43_DataGenerator/xxx.csv", true)
	if err != nil {
		t.Error("GetSeriesFromFile")
	}

	points, err := SeriesToPoints(series[0], series[1])
	if err != nil {
		t.Error("SeriesToPoints")
	}

	log.Println("Line Points: ", len(*points))
	visva := NewVisva(*points)
	res := 0.95
	simplx := visva.Simplify(res)

	l1 := geom.NewLineString(geom.Coordinates(*points))
	log.Println("Line Points l1: ", len(l1.Coordinates.Pnts))
	log.Println(l1.WKT())

	l2 := geom.NewLineString(geom.Coordinates(simplx))
	log.Println("Line Points l2: ", len(l2.Coordinates.Pnts))
	log.Println(l2.WKT())
}
