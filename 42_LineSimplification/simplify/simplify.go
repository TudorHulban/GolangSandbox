package simplify

import (
	"errors"
)

// Floats takes an array of points and returns the simplified version as per DouglasPeucker alghoritm
func Floats(points [][]float64, tolerance float64, highestQuality bool) ([][]float64, error) {
	if len(points) < 2 {
		return nil, errors.New("points array to small")
	}
	sqTolerance := tolerance * tolerance

	if highestQuality {
		return simplifyDouglasPeucker(points, sqTolerance), nil
	}
	simplePoints, err := simplifyRadialDistance(points, sqTolerance)
	if err != nil {
		return nil, err
	}
	return simplifyDouglasPeucker(simplePoints, sqTolerance), nil
}

func comparePoints(p1, p2 []float64) bool {
	if p1[0] == p2[0] && p1[1] == p2[1] {
		return true
	}
	return false
}
