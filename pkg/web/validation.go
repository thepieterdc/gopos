package web

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// Validator default request validator
type Validator struct {
	Validator *validator.Validate
}

func (v Validator) Validate(i interface{}) error {
	// Validate the input struct.
	if err := v.Validator.Struct(i); err != nil {
		// Convert the error to a validation error.
		validationError := err.(validator.ValidationErrors)[0]

		// Handle missing fields.
		if validationError.Tag() == "required" {
			field := strings.ToLower(validationError.Field())
			return echo.NewHTTPError(
				http.StatusUnprocessableEntity,
				fmt.Sprintf("required field \"%s\" was not found.", field),
			)
		}

		// Default error handler.
		return err
	}

	// Everything OK.
	return nil
}

// ParseAndValidate parses the request and applies data validation.
func ParseAndValidate(ctx *echo.Context, target interface{}) error {
	// Parse the arguments.
	if err := (*ctx).Bind(target); err != nil {
		// Parsing has failed.
		(*ctx).Logger().Error(err)
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"invalid parameter type was passed",
		)
	}

	// Validate the arguments.
	if err := (*ctx).Validate(target); err != nil {
		return err
	}

	// Input has been bound to the target and validation passed.
	return nil
}
