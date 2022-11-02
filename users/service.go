package users

import (
	"context"

	"github.com/vaibhavchalse99/db"
	"go.uber.org/zap"
)

type Service interface {
	Create(ctx context.Context, req createRequest) (err error)
	List(ctx context.Context) (response listResponse, err error)
	Login(ctx context.Context, req userCredentials) (response tokenResponse, err error)
	GetById(ctx context.Context, userId string) (user User, err error)
	UpdateById(ctx context.Context, req updateRequest, userId string) (user User, err error)
}

type userService struct {
	store  db.UserStorer
	logger *zap.SugaredLogger
}

func (us *userService) Create(ctx context.Context, req createRequest) (err error) {
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

func (us *userService) List(ctx context.Context) (response listResponse, err error) {
	dbUsers, err := us.store.GetUsers(ctx)

	if err != nil {
		if err == db.ErrUserNotExist {
			us.logger.Error("user not present", "err", err.Error())
			return response, ErrUserNotExist
		}
		us.logger.Error("Error listing users", "users", err.Error())
		return
	}
	for _, dbUser := range dbUsers {
		user := mapUserData(dbUser)
		response.Users = append(response.Users, user)
	}
	return
}

func (us *userService) GetById(ctx context.Context, userId string) (user User, err error) {
	if userId == "" {
		us.logger.Error("Invalid request for UserData", "msg", err.Error(), "user", userId)
		return user, errEmptyID
	}
	dbUser, err := us.store.GetUserDetailsById(ctx, userId)

	if err != nil {
		if err == db.ErrUserNotExist {
			us.logger.Error("user not present", "err", err.Error())
			return user, ErrUserNotExist
		}
		us.logger.Error("Error Getting user by Id", "user", err.Error())
		return
	}
	user = mapUserData(dbUser)
	return
}

func (us *userService) UpdateById(ctx context.Context, req updateRequest, userId string) (user User, err error) {
	err = req.Validate()
	if err != nil {
		us.logger.Error("Invalid request for user updation", "msg", err.Error(), "user", req)
		return user, errInvalidRequest
	}

	dbUser, err := us.store.UpdateUserDetailsById(ctx, userId, req.Name, req.Password)

	if err != nil {
		us.logger.Error("Error while updating the user", "user", err.Error())
		return
	}

	user = mapUserData(dbUser)
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
		return response, ErrUserNotExist
	}

	token, err := createToken(dbUser.ID, dbUser.Role)

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
