// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP
// Package gerror contains ...
package gerror

import "fmt"

// ErrorCode service error constants
type ErrorCode string

const (
	// InternalError ...
	InternalError ErrorCode = "Internal Error"
)

func (e ErrorCode) String() string {
	return string(e)
}

// GetErrorType ...
func GetErrorType(err error) ErrorCode {
	gerr, ok := err.(Gerror)
	if ok {
		return gerr.Tag().(ErrorCode)
	}
	return InternalError
}

// GetErrorMessage ...
func GetErrorMessage(err error) string {
	if gerr, ok := err.(Gerror); ok {
		if cause := gerr.Cause(); cause != nil {
			return fmt.Sprintf("%s: %s", gerr.Tag(), GetErrorMessage(cause))
		}
		return fmt.Sprintf("%s: %s", gerr.Tag(), gerr.Message())
	}
	return err.Error()
}
