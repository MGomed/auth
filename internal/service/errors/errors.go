package errors

import "errors"

// Token errors
var (
	ErrInvalidToken = errors.New("invalid token claims")
	ErrAccessDenied = errors.New("access denied")
)
