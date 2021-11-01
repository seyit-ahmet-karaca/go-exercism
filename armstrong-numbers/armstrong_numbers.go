package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	nStr := strconv.Itoa(n)
	nLen := len(nStr)
	nCopy := n
	sum := 0

	for nCopy != 0 {
		lastDigit := nCopy % 10
		nCopy /= 10
		pow := math.Pow(float64(lastDigit),float64(nLen))
		sum += int(pow)
	}
	return sum == n
}
