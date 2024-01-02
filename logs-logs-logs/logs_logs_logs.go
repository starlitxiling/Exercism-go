package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	// panic("Please implement the Application() function")
	apps := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001F50D': "search",
		'\u2600':     "weather",
	}

	for _, char := range log {
		if app, exists := apps[char]; exists {
			return app
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// panic("Please implement the Replace() function")

	// return strings.ReplaceAll(log, string(oldRune), string(newRune))

	var result strings.Builder
	for _, ch := range log {
		if ch == oldRune {
			result.WriteRune(newRune)
		} else {
			result.WriteRune(ch)
		}
	}
	return result.String()

}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// panic("Please implement the WithinLimit() function")
	numberOfRunes := utf8.RuneCountInString(log)

	if numberOfRunes > limit {
		return false
	}
	return true
}
