package models

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your item already exist")
	ErrBadParamInput       = errors.New("Given Param is not valid")
)
