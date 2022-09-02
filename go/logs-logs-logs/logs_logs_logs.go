package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	values := map[int]string{}
	for _, char := range log {
		v := fmt.Sprintf("%c", char)
		if v == "❗" || v == "🔍" || v == "☀" {
			values[len(values)] = v
		}
	}

	switch values[0] {
	case "❗":
		return "recommendation"
	case "🔍":
		return "search"
	case "☀":
		return "weather"
	default:
		return "default"
	}
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := strings.ReplaceAll(log, fmt.Sprintf("%c", oldRune), fmt.Sprintf("%c", newRune))
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
