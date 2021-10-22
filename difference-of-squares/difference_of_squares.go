package diffsquares

import (
	"math"
)

func SquareOfSum(number int) int {
	total := 0
	for i := 1 ; i <= number; i++ {
		total = total + i
	}

	return int(math.Pow(float64(total),2.0))
}

func SumOfSquares(number int) int {
	total := 0
	for i := 0 ; i <= number ; i++ {
		total = total + int(math.Pow(float64(i), 2.0))
	}

	return total
}

func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
