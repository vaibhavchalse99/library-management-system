package books

import "errors"

var (
	errEmptyID         = errors.New("Book id must be present")
	errEmptyName       = errors.New("Book name must be present")
	errEmptyAuthor     = errors.New("Book author must be present")
	errEmptyPrice      = errors.New("Book price must be present")
	errInvalidBookId   = errors.New("Invalid book id")
	errEmptyBookNumber = errors.New("Empty book number")
	errEmptyBookRef    = errors.New("Book reference should present")
)
