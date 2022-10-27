package books

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type createBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

type updateBookRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

type createBookCopy struct {
	ISBN   string `json:"isbn"`
	BookId string `json:"bookId"`
}

type deleteBookCopy struct {
	ISBN string `json:"isbn"`
}

type ISBNResponse struct {
	ISBN string `json:"isbn"`
}

func (req createBookRequest) Validate() (err error) {
	if req.Name == "" {
		return errEmptyName
	}
	if req.Author == "" {
		return errEmptyAuthor
	}
	if req.Price == 0 {
		return errEmptyPrice
	}
	return
}

func (req updateBookRequest) Validate() (err error) {
	if req.ID == "" {
		return errEmptyID
	}
	if req.Name == "" {
		return errEmptyName
	}
	if req.Author == "" {
		return errEmptyAuthor
	}
	if req.Price == 0 {
		return errEmptyPrice
	}
	return
}

func (req createBookCopy) Validate() (err error) {
	if req.ISBN == "" {
		return errEmptyBookNumber
	}
	if req.BookId == "" {
		return errEmptyBookRef
	}
	return
}

func (req deleteBookCopy) Validate() (err error) {
	if req.ISBN == "" {
		return errEmptyBookNumber
	}
	return
}
