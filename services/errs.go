package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase amont could not be zero")
	ErrRepository = errors.New("Repository Error")
)
