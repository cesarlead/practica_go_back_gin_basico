package domain

import "errors"

var (
	ErrInvalidID    = errors.New("id must be greater than 0")
	ErrInvalidName  = errors.New("name must be at least 5 characters long")
	ErrInvalidEmail = errors.New("invalid email format")
	ErrUserNotFound = errors.New("user not found")
)
