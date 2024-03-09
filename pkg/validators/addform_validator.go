package validators

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator"
)

func ValidateCyrillicOrLatinAndAscii(fl validator.FieldLevel) bool {
	projectName := fl.Field().String()

	for _, r := range projectName {
		if !unicode.IsLetter(r) && !unicode.Is(unicode.Space, r) && (r > 126 || r < 32) {
			fmt.Println("HERE")
			return false
		}

	}

	return true
}
