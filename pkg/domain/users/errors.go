package users

import "errors"

var (
	ErrInvalidUser     = errors.New("invalid user")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidID       = errors.New("invalid id")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrEmailNotFound   = errors.New("email not found")
)
