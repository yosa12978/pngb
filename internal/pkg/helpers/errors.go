package helpers

import "errors"

var (
	ErrBadRequest = errors.New("error: bad request")
	ErrNotFound   = errors.New("error: not found")
)
