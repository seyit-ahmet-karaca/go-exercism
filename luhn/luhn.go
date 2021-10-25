package luhn

import (
	"strconv"
	"strings"
)

//Valid checks card number according to luhn algorithm
func Valid(word string) bool {
	word = strings.ReplaceAll(word, " ", "")

	if len(word) <= 1 {
		return false
	}

	var sumOfDigits int64
	isDoubleIteration := false

	for i := len(word) - 1; i >= 0; i-- {
		intVal, err := strconv.ParseInt(string(word[i]), 10, 8)
		if err != nil {
			return false
		}
		if isDoubleIteration {
			intVal = intVal * 2
			if intVal > 9 {
				intVal = intVal - 9
			}
		}
		isDoubleIteration = !isDoubleIteration
		sumOfDigits = sumOfDigits + intVal
	}
	return sumOfDigits%10 == 0
}
