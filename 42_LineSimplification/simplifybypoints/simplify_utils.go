package simplify

func pointsSqSegDist(p, p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	var po Point
	if dx != 0 || dy != 0 {
		t := ((p.X-p1.X)*dx + (p.Y-p1.Y)*dy) / (dx*dx + dy*dy)
		if t > 1 {
			po = p2
		} else if t > 0 {
			po.X = p1.X + dx*t
			po.Y = p1.Y + dy*t
		}
	}
	dx = p.X - po.X
	dy = p.Y - po.Y
	return dx*dx + dy*dy
}

// radialDistance computes distances between provided points.
func radialDistance(pPoints []Point, pTolerance float64) ([]Point, error) {
	prevPoint := pPoints[0]
	result := []Point{prevPoint}
	var currentPoint Point

	for _, currentPoint = range pPoints {
		delta := pointSquareDistance(currentPoint, prevPoint)
		if delta > pTolerance {
			result = append(result, currentPoint)
			prevPoint = currentPoint
		}
	}
	if prevPoint != currentPoint {
		result = append(result, currentPoint)
	}
	return result, nil
}

func pointSquareDistance(pPoint1, pPoint2 Point) float64 {
	dx := pPoint1.X - pPoint2.X
	dy := pPoint1.Y - pPoint2.Y
	return dx*dx + dy*dy
}

// pointsDouglasPeucker is the actual alghorytm.
func pointsDouglasPeucker(pPoints []Point, sqTolerance float64) []Point {
	markers := make([]int, len(pPoints))
	firstPos := 0
	lastPos := len(pPoints) - 1

	var stackPoints []int
	result := []Point{}
	maxSqDist, sqDist := float64(0), float64(0)
	markers[firstPos], markers[lastPos] = 1, 1

	pop := func() int {
		if len(stackPoints) == 0 {
			return 0
		}
		result := stackPoints[len(stackPoints)-1]
		stackPoints = stackPoints[0 : len(stackPoints)-1]
		return result

	}
	index := 0
	for lastPos > 0 {
		maxSqDist = 0
		for i := firstPos + 1; i < lastPos; i++ {
			sqDist = pointsSqSegDist(pPoints[i], pPoints[firstPos], pPoints[lastPos])
			if sqDist > maxSqDist {
				index = i
				maxSqDist = sqDist
			}
		}
		if maxSqDist > sqTolerance {
			markers[index] = 1
			stackPoints = append(stackPoints, firstPos, index, index, lastPos)
		}
		lastPos = pop()
		firstPos = pop()
	}
	for i := 0; i < len(pPoints); i++ {
		if i < len(markers) && markers[i] > 0 {
			result = append(result, pPoints[i])
		}
	}
	return result
}
