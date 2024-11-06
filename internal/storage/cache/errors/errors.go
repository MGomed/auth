package errors

import "errors"

var ErrUserNotPresent = errors.New("user with such id is not present in cache")