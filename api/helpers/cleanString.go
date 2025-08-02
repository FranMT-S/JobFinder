package helpers

import (
	"regexp"
	"strings"
	"unicode"
)

// isAllowedRune is a function that returns true if the rune is allowed to be in the string
// it is used to filter out emojis and other special characters
func isAllowedRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsPunct(r) || unicode.IsSpace(r) || r == '$'
}

// CleanEmojiFromString is a function that removes emojis and trim the string
func CleanEmojiAndTrimString(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if isAllowedRune(r) {
			builder.WriteRune(r)
		}
	}

	return strings.TrimSpace(builder.String())
}

// CleanOnlyNumbersFromString is a function that removes all non-numeric characters from the string
func CleanOnlyNumbersFromString(s string) string {
	re := regexp.MustCompile(`[^0-9.]`)
	return re.ReplaceAllString(s, "")
}

// SanitizeParam is a function that sanitizes a string for use as a parameter
func SanitizeParam(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]`)
	s = re.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	return s
}
