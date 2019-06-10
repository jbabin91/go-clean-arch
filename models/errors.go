package models

import "errors"

var (
	// ErrInternalServerError is used when throwing an internal error
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFoundError is used when throwing an error when an item is not found
	ErrNotFoundError = errors.New("Your requested Item was not found")
	// ErrConflictError is used when throwing an error when an item already exists
	ErrConflictError = errors.New("You Item already exists")
)