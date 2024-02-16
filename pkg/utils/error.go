package utils

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrAlreadyExist = errors.New("already exists")
	// Add more common errors as needed
)
