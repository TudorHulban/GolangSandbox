/*
Package simplify provides line simplification alghorytm.
Tested tolerance values up to 0.18.
*/
package simplify

import (
	"errors"
)

// Point structure resembles chart point.
type Point struct {
	X float64
	Y float64
}

// Points takes an array of points and returns the simplified version as per DouglasPeucker alghoritm
func Points(points []Point, tolerance float64, highestQuality bool) ([]Point, error) {
	if len(points) < 2 {
		return nil, errors.New("points array to small")
	}
	if highestQuality {
		return pointsDouglasPeucker(points, tolerance), nil
	}

	simplePoints, err := radialDistance(points, tolerance)
	if err != nil {
		return nil, err
	}
	return pointsDouglasPeucker(simplePoints, tolerance), nil
}
