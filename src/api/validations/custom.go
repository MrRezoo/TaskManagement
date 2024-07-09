package validations

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			validationErrors = append(validationErrors, ValidationError{
				Property: err.Field(),
				Tag:      err.Tag(),
				Value:    err.Param(),
				Message:  err.Error(),
			})
		}
	}
	return &validationErrors
}

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
