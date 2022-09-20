package users

import (
	"time"

	"github.com/google/uuid"
)

type listResponse struct {
	Users []User `json:"users"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type tokenResponse struct {
	Token string `json:"token"`
}
type createRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type userCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	if req.Password == "" {
		return errEmptyPassword
	}
	return
}
