package validation

import "github.com/go-playground/validator/v10"

type FieldErrors struct {
	Field string
	Error string
}

func ValidateStruct(user interface{}) []*FieldErrors {
	var errors []*FieldErrors
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element FieldErrors
			element.Field = err.StructField()
			element.Error = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}
