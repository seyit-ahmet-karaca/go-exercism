package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate returns givens string acronym form.
func Abbreviate(s string) string {
	var acronym string
	reg, err := regexp.Compile("[^a-zA-Z0-9 ']+")
	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, " ")
	fields := strings.Fields(strings.ToUpper(s))
	for _,word := range fields{
		acronym += string(word[0])
	}
	return acronym
}
