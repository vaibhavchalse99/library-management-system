package books

import "errors"

var (
	errEmptyName     = errors.New("Book name must be present")
	errEmptyAuthor   = errors.New("Book author must be present")
	errEmptyPrice    = errors.New("Book price must be present")
	errInvalidBookId = errors.New("Invalid book id")
)
