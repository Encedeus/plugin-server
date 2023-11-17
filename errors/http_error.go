package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type HTTPError struct {
	code    int
	message string
}

func (e HTTPError) Error() string {
	return e.message
}

func NewHttpError(message string, code int) HTTPError {
	ve := HTTPError{
		message: fmt.Sprintf("%s,%d", message, code),
	}

	return ve
}

func IsHTTPError(err error) bool {
	if errors.As(err, &HTTPError{}) {
		return true
	}

	return false
}

func ErrorToHttpError(err error) (*HTTPError, error) {
	if !IsHTTPError(err) {
		return nil, errors.ErrUnsupported
	}

	a := strings.Split(err.Error(), ",")
	parseInt, err := strconv.Atoi(a[1])
	if err != nil {
		return nil, err
	}

	return &HTTPError{
		parseInt,
		a[0],
	}, nil
}
