package cerror

import (
	"errors"
)

const (
	InternalComplexErrorType ComplexErrorType = iota
)

type (
	ComplexError struct {
		err     error
		errType ComplexErrorType
	}

	ComplexErrorType uint8
)

func (ce *ComplexError) Error() string {
	return ce.err.Error()
}

func (ce *ComplexError) ErrorCode() int {
	return 1000
}

func (ce *ComplexError) ErrorType() string {
	switch ce.errType {

	case InternalComplexErrorType:
		return "internal"

	default:
		return "unsupported"
	}
}

func New(text string, errType ComplexErrorType) error {
	return &ComplexError{
		err:     errors.New(text),
		errType: errType,
	}
}

func Wrap(err error, errType ComplexErrorType) error {
	return &ComplexError{
		err:     err,
		errType: errType,
	}
}
