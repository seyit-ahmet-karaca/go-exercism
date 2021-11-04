package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	var sum int
	if isbn[9] == 'X' {
		sum = 10
	} else {
		if num, err := strconv.Atoi(string(rune(isbn[9]))); err != nil {
			return false
		} else {
			sum = num
		}
	}
	for i, v := range isbn[:len(isbn)-1] {
		if number, err := strconv.Atoi(string(v)); err != nil {
			return false
		} else {
			sum += number * (10 - i)
		}

	}

	return sum%11 == 0
}
