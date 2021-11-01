package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var anagrams []string
	subject = strings.ToLower(subject)
	for _, candidate := range candidates {
		isAnagram := true
		for _, sub := range subject {
			countInSubject := strings.Count(subject, string(sub))
			countInCandidate := strings.Count(strings.ToLower(candidate), string(sub))
			if countInCandidate != countInSubject || strings.ToLower(candidate) == subject || len(candidate) != len(subject) {
				isAnagram = false
				break
			}
		}
		if isAnagram {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
