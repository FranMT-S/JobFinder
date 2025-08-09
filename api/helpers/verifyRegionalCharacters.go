package helpers

import "unicode/utf8"

// verify if a rune is a regional indicator symbol (U+1F1E6 - U+1F1FF)
func IsRegionalCharacter(r rune) bool {
	return r >= 0x1F1E6 && r <= 0x1F1FF
}

// verify if the first two runes of a string are regional indicator symbols
func StartsWithRegionalCharacters(s string) bool {
	if len(s) < 8 { // Each symbol occupies 4 bytes in UTF-8, 2 symbols = 8 bytes minimum
		return false
	}

	r1, size1 := utf8.DecodeRuneInString(s)
	r2, _ := utf8.DecodeRuneInString(s[size1:])

	return IsRegionalCharacter(r1) && IsRegionalCharacter(r2)
}
