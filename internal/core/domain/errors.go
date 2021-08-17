package domain

import "errors"

var (
	errInvalidFrom = errors.New("from address is not valid")
	errInvalidTo = errors.New("to addresses is not valid")
)
