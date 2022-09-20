package users

import (
	"context"

	"github.com/vaibhavchalse99/db"
	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req createRequest) (err error)
	list(ctx context.Context) (response listResponse, err error)
	Login(ctx context.Context, req userCredentials) (response tokenResponse, err error)
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

func (us *userService) list(ctx context.Context) (response listResponse, err error) {
	dbUsers, err := us.store.GetUsers(ctx)
	if err == db.ErrUserNotExist {
		us.logger.Error("user not present", "err", err.Error())
		return response, errUserNotExist
	}
	if err != nil {
		us.logger.Error("Error listing users", "users", err.Error())
		return
	}
	for _, dbUser := range dbUsers {
		var user User
		user.ID = dbUser.ID
		user.Email = dbUser.Email
		user.Name = dbUser.Name
		user.Role = dbUser.Role
		user.Password = ""
		user.CreatedAt = dbUser.CreatedAt
		user.UpdatedAt = dbUser.UpdatedAt

		response.Users = append(response.Users, user)
	}
	return
}

func (us *userService) Login(ctx context.Context, req userCredentials) (response tokenResponse, err error) {
	err = req.Validate()
	if err != nil {
		us.logger.Error("Invalid request for login", "msg", err.Error(), "user", req)
		return
	}
	dbUser, err := us.store.GetUserDetails(ctx, req.Email, req.Password)

	if err != nil {
		us.logger.Error("User Not exist", "user", err.Error())
		return response, errUserNotExist
	}

	token, err := createToken(dbUser[0].ID)

	if err != nil {
		us.logger.Error("Error while creating the token", "user", err.Error())
		return
	}

	response.Token = token
	return
}

func NewService(dbUserStore db.UserStorer, logger *zap.SugaredLogger) Service {

	return &userService{
		store:  dbUserStore,
		logger: logger,
	}
}
