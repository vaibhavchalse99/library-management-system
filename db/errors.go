package db

import "errors"

var (
	ErrUserNotExist = errors.New("User doesn't exist in db")
)
