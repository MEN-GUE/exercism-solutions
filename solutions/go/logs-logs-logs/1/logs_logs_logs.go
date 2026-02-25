package logs
import "fmt"
import "unicode/utf8"


// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, char := range log {
        switch {
            case char == '❗':
            	return "recommendation"
            case char == '🔍':
            	return "search"
            case char == '☀':
            	return "weather"
    	}
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    runes := []rune(log)
	for i, char := range runes {
        if char == oldRune {
            runes[i] = newRune
        }
    }
    fmt.Println(runes)
    return string(runes)
}


// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
        return true
    }
    return false
}
