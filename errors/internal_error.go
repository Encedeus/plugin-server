package errors

import "errors"

type InternalError struct {
	message string
}

func (e InternalError) Error() string {
	return e.message
}

func NewInternalError(message string) InternalError {
	ve := InternalError{
		message: message,
	}

	return ve
}

func IsInternalError(err error) bool {
	if errors.As(err, &InternalError{}) {
		return true
	}

	return false
}
