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
