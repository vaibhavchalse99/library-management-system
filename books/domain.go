package books

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Price       int       `json:"price"`
	CopiesCount int       `json:"copies_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

type asssignBookCopy struct {
	BookCopyId string `json:"book_copy_id"`
	UserId     string `json:"user_id"`
	ReturnedAt string `json:"returned_at"`
}

type updateBookRecord struct {
	RecordId string `json:"record_id"`
}

type assignBookResponse struct {
	Message string `json:"message"`
}

type bookRecordDetails struct {
	Id             string    `json:"id"`
	BookISBN       string    `json:"book_isbn"`
	BookId         uuid.UUID `josn:"book_id"`
	BookName       string    `json:"book_name"`
	BookAuthor     string    `json:"book_author"`
	BookPrice      int       `json:"book_price"`
	UserId         uuid.UUID `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserEmail      string    `json:"user_email"`
	BookIssuedAt   time.Time `json:"book_issued_at"`
	BookReturnedAt time.Time `json:"book_returned_at"`
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

func (req asssignBookCopy) Validate() (err error) {

	if req.UserId == "" {
		return errEmptyUserId
	}
	if req.BookCopyId == "" {
		return errEmptyBookNumber
	}
	if req.ReturnedAt == "" {
		return errEmptyReturnedAt
	}
	return
}

func (req updateBookRecord) Validate() (err error) {
	if req.RecordId == "" {
		return errEmptyRecordId
	}
	return
}
