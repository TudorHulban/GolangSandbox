package simplify

import (
	"errors"
	"strconv"
)

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	if len(*s) > 0 {
		result := (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
		return result
	}
	return 0
}

func getSquareDistance(p1, p2 []float64) (float64, error) {
	if (len(p1) != 2) && (len(p2) != 2) {
		return 0, errors.New("problems with points: " + "p1 (len): " + strconv.FormatInt(int64(len(p1)), 10) + "p2 (len): " + strconv.FormatInt(int64(len(p2)), 10))
	}
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	return dx*dx + dy*dy, nil
}

func getSqSegDist(p, p1, p2 []float64) float64 {
	x := p1[0]
	y := p1[1]

	dx := p2[0] - x
	dy := p2[1] - y

	if dx != 0 || dy != 0 {
		t := ((p[0]-x)*dx + (p[1]-y)*dy) / (dx*dx + dy*dy)
		if t > 1 {
			x = p2[0]
			y = p2[1]
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}
	dx = p[0] - x
	dy = p[1] - y
	return dx*dx + dy*dy
}

func simplifyRadialDistance(points [][]float64, sqTolerance float64) ([][]float64, error) {
	prevPoint := points[0]
	result := [][]float64{prevPoint}
	var currentPoint []float64

	for i := 1; i < len(points); i++ {
		currentPoint = points[i]

		delta, err := getSquareDistance(currentPoint, prevPoint)
		if err != nil {
			return nil, err
		}
		if delta > sqTolerance {
			result = append(result, currentPoint)
			prevPoint = currentPoint
		}
	}
	if !comparePoints(prevPoint, currentPoint) {
		result = append(result, currentPoint)
	}
	return result, nil
}

func simplifyDouglasPeucker(pPoints [][]float64, sqTolerance float64) [][]float64 {
	markers := make([]int, len(pPoints))
	firstPos := 0
	lastPos := len(pPoints) - 1

	stackPoints := stack{}
	result := [][]float64{}
	maxSqDist, sqDist := float64(0), float64(0)
	markers[firstPos], markers[lastPos] = 1, 1

	index := 0
	for lastPos > 0 {
		maxSqDist = 0
		for i := firstPos + 1; i < lastPos; i++ {
			sqDist = getSqSegDist(pPoints[i], pPoints[firstPos], pPoints[lastPos])
			if sqDist > maxSqDist {
				index = i
				maxSqDist = sqDist
			}
		}
		if maxSqDist > sqTolerance {
			markers[index] = 1
			stackPoints.Push(firstPos)
			stackPoints.Push(index)
			stackPoints.Push(index)
			stackPoints.Push(lastPos)
		}
		lastPos = stackPoints.Pop()
		firstPos = stackPoints.Pop()
	}
	for i := 0; i < len(pPoints); i++ {
		if checkArrIndex(markers, i) {
			result = append(result, pPoints[i])
		}
	}
	return result
}

func checkArrIndex(arr []int, index int) bool {
	if index < len(arr) && index >= 0 {
		if arr[index] > 0 {
			return true
		}
	}
	return false
}
