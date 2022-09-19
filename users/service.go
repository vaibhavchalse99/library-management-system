package users

import (
	"context"

	"github.com/vaibhavchalse99/db"
	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req createRequest) (err error)
}

type userService struct {
	store  db.UserStorer
	logger *zap.SugaredLogger
}

func (us *userService) create(ctx context.Context, req createRequest) (err error) {
	err = req.Validate()

	if err != nil {
		us.logger.Errorw("Invalid request for user creation", "msg", err.Error(), "user", req)
		return
	}

	err = us.store.CreateUser(ctx, &db.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})

	if err != nil {
		us.logger.Errorw("Error creating user", "err", err.Error())
		return
	}
	return
}

func NewService(dbUserStore db.UserStorer, logger *zap.SugaredLogger) Service {

	return &userService{
		store:  dbUserStore,
		logger: logger,
	}
}
