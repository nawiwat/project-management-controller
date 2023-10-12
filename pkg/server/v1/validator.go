package server

import (
	"app-controller/pkg/errors"
	"net/http"

	"github.com/go-playground/validator"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		fields := map[string]interface{}{}
		for _, vErr := range err.(validator.ValidationErrors) {
			fields[vErr.Field()] = vErr.Tag()
		}
		return &errors.AppError{
			HTTPCode: http.StatusBadRequest,
			Code:     errors.ErrCodeValidationFail,
			Fields:   fields,
		}
	}
	return nil
}
