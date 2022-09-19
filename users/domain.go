package users

import "github.com/vaibhavchalse99/db"

type listResponse struct {
	Users []db.User `json:"users"`
}

type createRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
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
