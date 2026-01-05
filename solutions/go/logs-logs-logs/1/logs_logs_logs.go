package logs
import (
    "strings"
    "unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {

    for _, char := range log {
        switch char {
        case '\u2757':
            return "recommendation"
        case '\U0001F50D':
            return "search"
        case '\u2600':
            return "weather"
        }
    }
    return "default"
}
 
// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    oldRuneStr := string(oldRune)
    newRuneStr := string(newRune)
    
	result := strings.ReplaceAll(log, oldRuneStr, newRuneStr);
    
    return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	logLength := utf8.RuneCountInString(log)

    return logLength <= limit
}
