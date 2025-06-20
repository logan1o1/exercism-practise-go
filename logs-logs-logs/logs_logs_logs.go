package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, l := range log {
		switch l {
		case rune(0x2757):
			return "recommendation"
		case rune(0x1F50D):
			return "search"
		case rune(0x2600):
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog []rune
	for _, l := range log {
		if l != oldRune {
			newLog = append(newLog, l)
		} else {
			newLog = append(newLog, newRune)
		}
	}
	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
