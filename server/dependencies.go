package server

import (
	"github.com/vaibhavchalse99/app"
	"github.com/vaibhavchalse99/db"
	"github.com/vaibhavchalse99/users"
)

type dependencies struct {
	UserService users.Service
}

func initDependencies() (dependencies, error) {
	appDb := app.GetDB()
	logger := app.GetLogger()

	dbUserStore := db.NewUserStorer(appDb)

	userService := users.NewService(dbUserStore, logger)

	return dependencies{
		UserService: userService,
	}, nil

}
