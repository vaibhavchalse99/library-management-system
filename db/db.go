package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ctxKey int

const (
	dbKey          ctxKey = 0
	defaultTimeout        = 1 * time.Second
)

type UserStorer interface {
	CreateUser(ctx context.Context, user *User) (err error)
	GetUsers(ctx context.Context) (users []User, err error)
	GetUserDetails(ctx context.Context, email string, password string) (user User, err error)
	GetUserDetailsById(ctx context.Context, userId string) (users User, err error)
	UpdateUserDetailsById(ctx context.Context, userId string, name string, password string) (user User, err error)
}

type userStore struct {
	db *sqlx.DB
}

func Transact(ctx context.Context, dbx *sqlx.DB, opts *sql.TxOptions, txFunc func(context.Context) error) (err error) {
	tx, err := dbx.BeginTxx(ctx, opts)
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = errors.WithStack(p)
			default:
				err = errors.Errorf("%s", p)
			}
		}
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				err = errors.WithStack(e)
			}
			return
		}
		err = errors.WithStack(tx.Commit())
	}()

	ctxWithTx := newContext(ctx, tx)
	err = WithDefaultTimeout(ctxWithTx, txFunc)
	return err
}

func newContext(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, dbKey, tx)
}

func WithTimeout(ctx context.Context, timeout time.Duration, op func(ctx context.Context) error) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return op(ctxWithTimeout)
}

func WithDefaultTimeout(ctx context.Context, op func(ctx context.Context) error) (err error) {
	return WithTimeout(ctx, defaultTimeout, op)
}

func NewUserStorer(d *sqlx.DB) UserStorer {
	return &userStore{
		db: d,
	}
}

type bookStore struct {
	db *sqlx.DB
}

type BookStorer interface {
	CreateBook(ctx context.Context, name string, author string, price int) (book Book, err error)
	BookList(ctx context.Context) (books []Book, err error)
	GetBookById(ctx context.Context, bookId string) (book Book, err error)
	UpdateBook(ctx context.Context, id, author, name string, price int) (book Book, err error)
	AddBookcopy(ctx context.Context, isbn, bookId string) (bookIsbn string, err error)
	RemoveBookcopy(ctx context.Context, isbn string) (bookIsbn string, err error)
	AssignBook(ctx context.Context, bookCopyId, userId string, returnedAt time.Time) (err error)
	GetBookId(ctx context.Context, book_copy_id string) (bookId string, err error)
	GetAllIssuedBookIds(ctx context.Context, userId string) (bookIds []string, err error)
	GetRecordsInfoByIsbnNumber(ctx context.Context, isbn string) (bookRecord bookRecordDetails, err error)
}

func NewBookStorer(d *sqlx.DB) BookStorer {
	return &bookStore{
		db: d,
	}
}
