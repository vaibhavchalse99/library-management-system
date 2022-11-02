package books

import "errors"

var (
	errEmptyID            = errors.New("book id must be present")
	errEmptyName          = errors.New("book name must be present")
	errEmptyAuthor        = errors.New("book author must be present")
	errEmptyPrice         = errors.New("book price must be present")
	errInvalidBookId      = errors.New("invalid book id")
	errEmptyBookNumber    = errors.New("empty book number")
	errEmptyBookRef       = errors.New("book reference should present")
	errEmptyUserId        = errors.New("user id must be present")
	errEmptyReturnedAt    = errors.New("book returned data must be present")
	errEBookAlreadyIssued = errors.New("this book is already issued by the user")
	errEmptyRecordId      = errors.New("record id must be present")
)
