package pangram

import "strings"

func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)

	for i := 'a' ; i <= 'z' ; i++ {
		if strings.Count(sentence, string(i)) < 1 {
			return false
		}
	}
	return true
}