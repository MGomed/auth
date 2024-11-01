package errors

import "errors"

var (
	ErrQueryBuild   = errors.New("failed to build query")
	ErrQueryExecute = errors.New("failed to execute query")
)
