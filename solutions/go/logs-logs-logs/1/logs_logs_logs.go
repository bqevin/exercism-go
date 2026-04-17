package logs

import (
    "unicode/utf8"
    "strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
    for len(log) > 0 {
		r, size := utf8.DecodeRuneInString(log)
		if r == '❗' {
            return "recommendation";
        } else if r == '🔍' {
            return "search";
        } else if r ==  '☀' {
            return "weather";
        }
		log = log[size:]
	}
    return "default";
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Replace(log, string(oldRune), string(newRune), -1);
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
