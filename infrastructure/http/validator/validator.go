package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	val := validator.New()

	return &Validator{
		validator: val,
	}

}

func (v *Validator) Validate(data interface{}) error {
	if err := v.validator.Struct(data); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}
