package darts

import "math"

const (
	tenPoint = 10
	fivePoint = 5
	onePoint = 1
	zeroPoint = 0
)

const (
	innerCircleRadius = 1
	middleCircleRadius = 5
	outerCircleRadius = 10
)

func Score(x, y float64) int {
	distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if distance <= innerCircleRadius {
		return tenPoint
	}
	if distance <= middleCircleRadius {
		return fivePoint
	}
	if distance <= outerCircleRadius {
		return onePoint
	}
	return zeroPoint

}
