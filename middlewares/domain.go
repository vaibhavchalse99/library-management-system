package middlewares

import "errors"

var (
	ErrTokenNotPresent = errors.New("Token is not present")
	ErrInvalidToken    = errors.New("Token is Invalid")
)
