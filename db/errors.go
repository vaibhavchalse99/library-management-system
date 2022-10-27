package db

import "errors"

var (
	ErrUserNotExist       = errors.New("User doesn't exist")
	ErrBooksNotExist      = errors.New("Books not exist")
	ErrSomethingWentWrong = errors.New("Something went wrong")
)
