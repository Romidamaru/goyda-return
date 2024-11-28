package utils

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// ValidateStruct validates the given struct and returns an error if validation fails.
func (v *Validator) ValidateStruct(s any) error {
	return v.validate.Struct(s)
}
