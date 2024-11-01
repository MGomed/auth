package errors

import "errors"

var (
	ErrNameLenInvalid   = errors.New("name's length should be between 2 and 32")
	ErrEmailInvalid     = errors.New("wrong email format")
	ErrPasswordMismatch = errors.New("password didn't match")
)
