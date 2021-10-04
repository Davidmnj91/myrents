package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func BirthDateValidator(fl validator.FieldLevel) bool {
	birthDate := fl.Field().String()
	re := regexp.MustCompile(`[\d]{4}-[\d]{2}-[\d]{2}`)

	return re.MatchString(birthDate)
}
