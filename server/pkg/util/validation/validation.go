package validation

import (
	"github.com/go-playground/validator/v10"
)

type FieldErrors struct {
	Field string
	Error string
}

type Validator interface {
	ValidateStruct(toValidate interface{}) []*FieldErrors
}

type validate struct {
	validator *validator.Validate
}

func NewValidator() (Validator, error) {
	v := validator.New()
	err := v.RegisterValidation("birthDate", BirthDateValidator)
	if err != nil {
		return nil, err
	}

	return &validate{validator: v}, nil
}

func (v *validate) ValidateStruct(toValidate interface{}) []*FieldErrors {
	var errors []*FieldErrors
	err := v.validator.Struct(toValidate)
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
