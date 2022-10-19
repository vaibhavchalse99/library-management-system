package server

import (
	"github.com/vaibhavchalse99/app"
	"github.com/vaibhavchalse99/books"
	"github.com/vaibhavchalse99/db"
	"github.com/vaibhavchalse99/users"
)

type dependencies struct {
	UserService users.Service
	BookService books.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()

	dbUserStore := db.NewUserStorer(appDB)
	dbBookStore := db.NewBookStorer(appDB)

	userService := users.NewService(dbUserStore, logger)
	bookService := books.NewService(dbBookStore, logger)

	return dependencies{
		UserService: userService,
		BookService: bookService,
	}, nil

}
