package users

import "errors"

var (
	errEmptyName      = errors.New("User name must be present")
	errEmptyID        = errors.New("User id must be present")
	errEmptyEmail     = errors.New("User email must be present")
	errEmptyPassword  = errors.New("User password must be present")
	errEmptyRole      = errors.New("User role must be present")
	ErrUserNotExist   = errors.New("User not exist")
	errInvalidRequest = errors.New("Request data not present")
)
