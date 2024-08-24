package internalErrors

import "errors"

var (
	DBErrDuplicateKey = errors.New("duplicate key error")
)
