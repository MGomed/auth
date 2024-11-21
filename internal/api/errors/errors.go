package errors

import "errors"

// ErrPasswordMismatch is Password and it's confirmation mismatch
var ErrPasswordMismatch = errors.New("password doesn't match")
