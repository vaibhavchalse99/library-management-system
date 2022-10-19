package books

import (
	"context"

	"github.com/vaibhavchalse99/db"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, req createBookRequest) (book Book, err error)
	List(ctx context.Context) (books []Book, err error)
	GetBookById(ctx context.Context, bookId string) (book Book, err error)
}

type bookService struct {
	store  db.BookStorer
	logger *zap.SugaredLogger
}

func (bs *bookService) Create(ctx context.Context, req createBookRequest) (book Book, err error) {
	err = req.Validate()
	if err != nil {
		bs.logger.Errorw("invalid request for book creatation", "msg", err.Error(), "req", req)
		return
	}

	dbBook, err := bs.store.CreateBook(ctx, req.Name, req.Author, req.Price)

	if err != nil {
		bs.logger.Errorw("Error while creating a book", "error", err.Error())
		return
	}
	book = Book(dbBook)
	return
}

func (bs *bookService) List(ctx context.Context) (books []Book, err error) {
	dbBooks, err := bs.store.BookList(ctx)
	if err != nil {
		bs.logger.Errorw("Error while fetching book list", "msg", err.Error())
		return
	}

	for _, dbBook := range dbBooks {
		book := Book(dbBook)
		books = append(books, book)
	}
	return
}

func (bs *bookService) GetBookById(ctx context.Context, bookId string) (book Book, err error) {
	if bookId == "" {
		bs.logger.Errorw("Book id is not present", "msg", "req", bookId)
		return book, errInvalidBookId
	}

	dbBook, err := bs.store.GetBookById(ctx, bookId)

	if err != nil {
		if err == db.ErrBooksNotExist {
			return book, db.ErrBooksNotExist
		}
		bs.logger.Errorw("Error while fetching the Book data", "msg", err.Error())
		return
	}
	book = Book(dbBook)
	return
}

func NewService(dbBookStore db.BookStorer, logger *zap.SugaredLogger) Service {
	return &bookService{
		store:  dbBookStore,
		logger: logger,
	}
}
