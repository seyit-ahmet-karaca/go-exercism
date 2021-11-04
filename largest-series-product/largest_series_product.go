package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) {
		return 0 ,errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return 0,errors.New("span must not be negative")
	}
	if len(digits) == 0 || span == 0{
		return 1 ,nil
	}
	digitsInt, err := mapToInt(digits)
	if err != nil {
		return 0,err
	}

	var total int
	i := 0
	for {
		subTotal := 1
		multiplied := false
		if i + span <= len(digitsInt) {
			for j:= i; j < i+span ; j++{
				subTotal *= digitsInt[j]
				multiplied = true
			}
		}
		if subTotal > total && multiplied{
			total = subTotal
		}
		i++
		if i == len(digitsInt) {
			break
		}
	}
	return total,nil
}

func mapToInt(digits string) ([]int,error) {
	var digitInt []int
	for _,v := range digits{
		if num, err := strconv.Atoi(string(v)); err == nil {
			digitInt = append(digitInt, num)
		} else {
			return nil, errors.New("digits input must only contain digits")
		}
	}
	return digitInt, nil
}