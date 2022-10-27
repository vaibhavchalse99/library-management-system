package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Author    string    `db:"author"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var (
	CreateBookQuery  = `INSERT INTO books(name, author, price, created_at, updated_at) VALUES($1,$2,$3,$4,$5) RETURNING *`
	GetBookListQuery = `SELECT * FROM books`
	GetBookByIdQuery = `SELECT * FROM books WHERE id = $1`
	UpdateBookQuery  = `UPDATE books SET author=$1, name=$2, price=$3, updated_at=$4 where ID=$5 RETURNING* `

	AddBookCopyQuery = `INSERT INTO book_copies(isbn, book_id, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING isbn`
)

func (d *bookStore) CreateBook(ctx context.Context, name string, author string, price int) (book Book, err error) {
	now := time.Now()
	err = d.db.GetContext(ctx, &book, CreateBookQuery, name, author, price, now, now)
	return
}

func (d *bookStore) BookList(ctx context.Context) (books []Book, err error) {
	err = d.db.SelectContext(ctx, &books, GetBookListQuery)
	if err == sql.ErrNoRows {
		return books, ErrBooksNotExist
	}
	return
}

func (d *bookStore) GetBookById(ctx context.Context, bookId string) (book Book, err error) {
	err = d.db.GetContext(ctx, &book, GetBookByIdQuery, bookId)
	if err == sql.ErrNoRows {
		return book, ErrBooksNotExist
	}
	return
}

func (d *bookStore) UpdateBook(ctx context.Context, id, author, name string, price int) (book Book, err error) {
	now := time.Now()
	err = d.db.GetContext(ctx, &book, UpdateBookQuery, author, name, price, now, id)
	if err == sql.ErrNoRows {
		return book, ErrBooksNotExist
	}
	return
}

func (d *bookStore) AddBookcopy(ctx context.Context, isbn, bookId string) (bookIsbn string, err error) {
	now := time.Now()
	err = d.db.GetContext(ctx, &bookIsbn, AddBookCopyQuery, isbn, bookId, now, now)
	if err != nil {
		return bookIsbn, err
	}
	return
}
