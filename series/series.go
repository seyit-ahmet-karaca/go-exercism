package series

import "strings"

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	if n == 1 {
		return strings.Split(s,"")
	}

	output := []string{}
	for i := 0 ; i <= len(s)-n; i++ {
		output = append(output, s[i:i+n])
	}

	return output
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int,s string) (first string, ok bool) {
	if n > len(s) || 0 >= n {
		return "", false
	}
	return s[:n], true
}