package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}
		count := strings.Count(strings.ToLower(word),string(char))
		if count > 1 {
			return false
		}
	}
	return true
}