package scrabble

import (
	"strings"
)

var pointsByLetter = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}

func Score(word string) int {
	total := 0
	for key, value := range pointsByLetter {
		for _, v := range value {
			count := strings.Count(strings.ToLower(word), strings.ToLower(v))
			total = total + key*count
		}
	}
	return total
}
