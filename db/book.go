package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Author      string    `db:"author"`
	Price       int       `db:"price"`
	CopiesCount int       `db:"copies_count"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

var (
	CreateBookQuery  = `INSERT INTO books(name, author, price, created_at, updated_at) VALUES($1,$2,$3,$4,$5) RETURNING *`
	GetBookListQuery = `SELECT books.*,(SELECT COUNT(*) as copies_count FROM book_copies WHERE book_copies.book_id = books.id) FROM books`
	GetBookByIdQuery = `SELECT * FROM books WHERE id = $1`
	UpdateBookQuery  = `UPDATE books SET author=$1, name=$2, price=$3, updated_at=$4 where ID=$5 RETURNING* `

	AddBookCopyQuery    = `INSERT INTO book_copies(isbn, book_id, created_at, updated_at) VALUES($1,$2,$3,$4) RETURNING isbn`
	DeleteBookCopyQuery = `DELETE FROM book_copies WHERE isbn=$1 RETURNING isbn`

	AssignBookQuery = `INSERT INTO records(book_copy_id, user_id, returned_at) VALUES($1,$2,$3)`

	GetBookIdQuery           = `SELECT book_id FROM book_copies where isbn=$1`
	GetAllIssuedBookIdsQuery = `SELECT bc.book_id FROM book_copies bc, records r WHERE r.book_copy_id = bc.isbn AND r.user_id = $1 AND $2 > r.issued_at AND $2 < r.returned_at`
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
	fmt.Println(books)
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

func (d *bookStore) RemoveBookcopy(ctx context.Context, isbn string) (bookIsbn string, err error) {
	err = d.db.GetContext(ctx, &bookIsbn, DeleteBookCopyQuery, isbn)
	if err != nil {
		return bookIsbn, err
	}
	return
}

func (d *bookStore) AssignBook(ctx context.Context, bookCopyId, userId string, returnedAt time.Time) (err error) {

	_, err = d.db.Query(AssignBookQuery, bookCopyId, userId, returnedAt)
	if err != nil {
		return err
	}
	return
}

func (d *bookStore) GetBookId(ctx context.Context, book_copy_id string) (bookId string, err error) {
	err = d.db.GetContext(ctx, &bookId, GetBookIdQuery, book_copy_id)
	if err == sql.ErrNoRows {
		return bookId, ErrBooksNotExist
	}
	return
}

func (d *bookStore) GetAllIssuedBookIds(ctx context.Context, userId string) (bookIds []string, err error) {
	err = d.db.SelectContext(ctx, &bookIds, GetAllIssuedBookIdsQuery, userId, time.Now())
	if err == sql.ErrNoRows {
		return bookIds, ErrBooksNotExist
	}
	return
}
