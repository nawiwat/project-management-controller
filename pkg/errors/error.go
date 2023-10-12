package errors

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

var _ error = &AppError{}

// AppError as standard error
type AppError struct {
	HTTPCode int         `json:"http_code"`
	Code     ErrCode     `json:"code"`
	Reason   string      `json:"reason,omitempty"`
	Fields   interface{} `json:"fields,omitempty"`
	Inner    error       `json:"inner,omitempty"`
}

// WithCode to set code
func (e *AppError) WithCode(code ErrCode) *AppError {
	e.Code = code
	return e
}

// WithReason to set reason
func (e *AppError) WithReason(reason string) *AppError {
	e.Reason = reason
	return e
}

// WithFields to set fields
func (e *AppError) WithFields(fields interface{}) *AppError {
	e.Fields = fields
	return e
}

// WithInner to set inner erro
func (e *AppError) WithInner(inner error) *AppError {
	e.Inner = inner
	return e
}

// New to create appError
func New(httpCode int, code ErrCode) *AppError {
	return &AppError{
		HTTPCode: httpCode,
		Code:     code,
	}
}

func (e *AppError) Error() string {
	raw, _ := json.Marshal(*e)
	err := "error: " + string(raw)
	if e.Inner != nil {
		err = "with inner: " + e.Inner.Error()
	}
	return err
}

// NewError to create app error
func NewError(httpCode int, code ErrCode, reason string, fields interface{}, inner error) error {
	return &AppError{
		HTTPCode: httpCode,
		Code:     code,
		Reason:   reason,
		Fields:   fields,
		Inner:    errors.WithStack(inner),
	}
}

// NewInternalError to internal error any error as standard
func NewInternalError(err error) error {
	return &AppError{
		HTTPCode: http.StatusInternalServerError,
		Code:     ErrCodeInternalError,
		Reason:   "",
		Fields:   nil,
		Inner:    err,
	}
}

// AssertAppError to assert any error as standard
func AssertAppError(err error) (*AppError, bool) {
	if appErr, ok := err.(*AppError); ok {
		return appErr, true
	}

	return &AppError{
		HTTPCode: http.StatusInternalServerError,
		Code:     ErrCodeInternalError,
		Reason:   "",
		Fields:   nil,
		Inner:    err,
	}, false
}
