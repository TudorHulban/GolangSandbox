package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/intdxdt/geom"
	"github.com/intdxdt/math"
)

type GeomPoint struct {
	geom.Point
	area float64
}

// Area - triangle area
func Area(a, b, c *geom.Point) float64 {
	return 0.5 * math.Abs((a[0]-c[0])*(b[1]-a[1])-(a[0]-b[0])*(c[1]-a[1]))
}

type Triangle struct {
	a, b, c *GeomPoint
	prev    *Triangle
	next    *Triangle
}

func NewTriangle(pts []*GeomPoint) *Triangle {
	var a, b, c = pts[0], pts[1], pts[2]

	b.area = Area(&a.Point, &b.Point, &c.Point)
	return &Triangle{a: a, b: b, c: c, prev: nil, next: nil}
}

func TriangleAreaCompare(t, o interface{}) int {
	dx := float64(t.(*Triangle).b.area - o.(*Triangle).b.area)

	if math.FloatEqual(dx, 0.0) {
		return 0
	}

	if dx < 0 {
		return -1
	}

	return 1
}

func (t *Triangle) String() string {
	return geom.NewPolygon(geom.Coordinates([]geom.Point{t.a.Point, t.b.Point, t.c.Point})).WKT()
}

func SeriesToPoints(seriesX, seriesY []float64) (*[]geom.Point, error) {
	if len(seriesX) != len(seriesY) {
		return nil, errors.New("Series are not of equal values")
	}
	noPoints := len(seriesX)

	instance := make([]geom.Point, noPoints)
	for i := 0; i < noPoints-1; i++ {
		instance[i] = geom.Point{seriesX[i], seriesY[i]}
	}
	return &instance, nil
}

func GetSeriesFromFile(filePath string, withHeader bool) ([][]float64, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	scanner1 := bufio.NewScanner(strings.NewReader(string(file)))
	var fields []string
	for scanner1.Scan() {
		fields = strings.Split(scanner1.Text(), ",")
	}
	noFields := len(fields)
	log.Println("noFields: ", noFields)
	result := make([][]float64, noFields)

	scanner2 := bufio.NewScanner(strings.NewReader(string(file)))
	if withHeader {
		scanner2.Scan()
	}

	for scanner2.Scan() {
		rowStrings := strings.Split(scanner2.Text(), ",")

		for k, v := range rowStrings {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				log.Println("string conv: ", err)
				return nil, err
			}
			result[k] = append(result[k], f)
		}
	}
	log.Println("len:", len(result), result[0], result[1])
	return result, err
}
