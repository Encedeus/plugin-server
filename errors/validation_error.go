package errors

import "errors"

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func NewValidationError(message string) ValidationError {
	ve := ValidationError{
		message: message,
	}

	return ve
}

func IsValidationError(err error) bool {
	if errors.As(err, &ValidationError{}) {
		return true
	}

	return false
}
