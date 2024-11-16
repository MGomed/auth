package errors

import "errors"

// MessageBus errors
var (
	ErrMarshal      = errors.New("failed to encode msg to json")
	ErrUnrecognized = errors.New("unrecognized message type")
)
