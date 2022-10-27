package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/vaibhavchalse99/db"
)

type listResponse struct {
	Users []User `json:"users"`
}

type response struct {
	User User `json:"user"`
}

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"password,omitempty"`
	Role      db.RoleValue `json:"role"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type tokenResponse struct {
	Token string `json:"token"`
}
type createRequest struct {
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Password string       `json:"password"`
	Role     db.RoleValue `json:"role"`
}

type updateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type userCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req updateRequest) Validate() (err error) {
	if req.Name == "" && req.Password == "" {
		return errInvalidRequest
	}
	return
}

func (req userCredentials) Validate() (err error) {
	if req.Email == "" {
		return errEmptyEmail
	}
	if req.Password == "" {
		return errEmptyPassword
	}
	return
}

func (req createRequest) Validate() (err error) {
	if req.Name == "" {
		return errEmptyName
	}
	if req.Email == "" {
		return errEmptyEmail
	}
	if req.Role == "" {
		return errEmptyRole
	}
	if err = req.Role.Validate(); err != nil {
		return
	}
	if req.Password == "" {
		return errEmptyPassword
	}
	return
}
