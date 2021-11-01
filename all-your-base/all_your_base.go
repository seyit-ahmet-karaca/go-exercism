package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	value, err := valueFromBase(inputBase, inputDigits)
	if err != nil {
		return nil, err
	}
	outputDigits, err := digitsFromValue(outputBase, value)
	if err != nil {
		return nil, err
	}
	return outputDigits, nil
}

func valueFromBase(base int, digits []int) (int, error) {
	if base < 2 {
		return 0, errors.New("input base must be >= 2")
	}
	var sum int
	for _, digit := range digits {
		if digit < 0 || base <= digit {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		sum *= base
		sum += digit
	}
	return sum, nil
}

func digitsFromValue(base int, value int) ([]int, error) {
	if base < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if value == 0 {
		return []int{0}, nil
	}
	var digits []int
	for value > 0 {
		ints := []int{value % base}
		digits = append(ints, digits...)
		value /= base
	}
	return digits, nil
}
