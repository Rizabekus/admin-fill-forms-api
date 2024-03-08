package validators

import (
	"unicode"

	"github.com/go-playground/validator"
)

func validateCyrillicOrLatinAndAscii(fl validator.FieldLevel) bool {
	projectName := fl.Field().String()
	hasCyrillic := false
	hasLatin := false

	for _, r := range projectName {
		if !unicode.IsLetter(r) && !unicode.Is(unicode.Space, r) && (r > 126 || r < 32) {
			return false
		}
		if unicode.Is(unicode.Cyrillic, r) {
			hasCyrillic = true
		}
		if unicode.Is(unicode.Latin, r) {
			hasLatin = true
		}
	}

	return hasCyrillic && hasLatin
}
