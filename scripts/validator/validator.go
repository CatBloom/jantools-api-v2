package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &customValidator{validator: validator.New()}
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
