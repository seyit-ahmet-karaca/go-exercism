package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation, _ := utf8.DecodeRuneInString("❗")
	search, _ := utf8.DecodeRuneInString("🔍")
	weather, _ := utf8.DecodeRuneInString("☀")

	fmt.Println(recommendation)
	utf8.DecodeRuneInString(log)
	for len(log) > 0 {
		r, size := utf8.DecodeRuneInString(log)
		if recommendation == r {
			return "recommendation"
		} else if weather == r {
			return "weather"
		} else if search == r {
			return "search"
		}
		log = log[size:]
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {

	return utf8.RuneCountInString(log) <= limit
}
