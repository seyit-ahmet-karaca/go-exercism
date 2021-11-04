package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	decrypted := ""
	decryptedArr := []string{}
	for _, v := range s {
		isDigit := unicode.IsDigit(v)

		if !unicode.IsLetter(v) && !isDigit {
			continue
		}
		if isDigit {
			decrypted += string(v)
		} else {
			lower := strings.ToLower(string(v))[0]
			newLetter := 25 - (lower - 'a')
			decrypted += string('a' + newLetter)
		}
	}

	for i := 0; i < len(decrypted); i += 5 {
		start, end := i, i+5
		if end > len(decrypted) {
			end = len(decrypted)
		}
		decryptedArr = append(decryptedArr, decrypted[start:end])
	}

	return strings.Join(decryptedArr, " ")
}
