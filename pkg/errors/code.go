package errors

// ErrCode to define error code type
type ErrCode string

const (
	// ErrCodeInternalError for unknow internal error
	ErrCodeInternalError ErrCode = "INTERNAL_ERROR"
	// ErrCodeValidationFail for validation fields
	ErrCodeValidationFail ErrCode = "VALIDATION_FAIL"
	ErrNotFound           ErrCode = "NOT_FOUND"
	ErrBadRequest         ErrCode = "BAD_REQUEST"
)
