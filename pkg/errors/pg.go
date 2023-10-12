package errors

import (
	"strings"
)

type PQErrCode string

const (
	UniqueViolationErr     PQErrCode = "23505"
	ForeignKeyViolationErr PQErrCode = "23503"
)

func IsErrorCode(err error, errcode PQErrCode) bool {
	if ok := strings.Contains(err.Error(), "SQLSTATE"); ok {
		var code string
		code = strings.Split(err.Error(), "SQLSTATE")[1]
		code = code[1 : len(code)-1]
		return code == string(errcode)
	}
	return false
}
