package errors

import "errors"

// ErrUserNotPresent indicates about user's existance in cache
var ErrUserNotPresent = errors.New("user with such id is not present in cache")
