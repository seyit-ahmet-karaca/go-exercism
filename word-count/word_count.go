package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(sentence string) Frequency {
	freq := Frequency{}
	match, _ := regexp.Compile("[^A-Za-z0-9']+")
	split := match.Split(sentence, -1)

	for _,token := range split {
		if token != "" {
			freq[strings.Trim(strings.ToLower(token),"'")]++
		}
	}

	return freq
}
