package bob

import (
	"regexp"
	"strings"
)

// Hey "listens" to it's input remark and responds based on what it hears.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if  remark == "" {
		return "Fine. Be that way!"
	}
	isLetters, _ := regexp.MatchString("[a-zA-Z]+", remark)

	yelling := strings.ToUpper(remark) == remark && isLetters
	question := strings.HasSuffix(remark, "?")
	if yelling && question {
		return "Calm down, I know what I'm doing!"
	}
	if yelling && !question {
		return "Whoa, chill out!"
	}
	if !yelling && question {
		return "Sure."
	}
	return "Whatever."
}