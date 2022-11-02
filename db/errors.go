package db

import "errors"

var (
	ErrUserNotExist          = errors.New("user doesn't exist")
	ErrBooksNotExist         = errors.New("books not exist")
	ErrSomethingWentWrong    = errors.New("something went wrong")
	ErrBooksNotAssigned      = errors.New("this book is currently not assigned to any user")
	ErrBooksCopyNotExist     = errors.New("book copy not exist")
	ErrBooksCopyNotAvailable = errors.New("book copy not available")
	ErrBookCopyAlreadyIssued = errors.New("books copy already issued")
)
