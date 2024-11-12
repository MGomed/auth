package errors

import "errors"

// Repo errors
var (
	ErrQueryBuild   = errors.New("failed to build query")
	ErrQueryExecute = errors.New("failed to execute query")
)
