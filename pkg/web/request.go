package web

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// Validator default request validator
type Validator struct {
	Validator *validator.Validate
}

func (v Validator) Validate(i interface{}) error {
	// Validate the input struct.
	if err := v.Validator.Struct(i); err != nil {
		return err
	}

	// Everything OK.
	return nil
}

// ParseAndValidate parses the request and applies data validation.
func ParseAndValidate(ctx *echo.Context, target interface{}) error {
	// Parse the arguments.
	if err := (*ctx).Bind(target); err != nil {
		return err
	}

	// Validate the arguments.
	if err := (*ctx).Validate(target); err != nil {
		return err
	}

	// Input has been bound to the target and validation passed.
	return nil
}
