package etl

import "strings"

func Transform(m map[int][]string) map[string]int {
	var newBrand = map[string]int{}

	for key,val := range m {
		for _, v := range val {
			newBrand[strings.ToLower(v)] = key
		}
	}
	return newBrand
}