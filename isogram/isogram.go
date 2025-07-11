package isogram

import (
	"slices"
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	run := []rune(word)
	for i, r := range run {
		if r == ' ' || r == '-' {
			continue
		}

		r = unicode.ToLower(r)

		if slices.Contains(run[i+1:], r) {
			return false
		}
	}
	return true
}
