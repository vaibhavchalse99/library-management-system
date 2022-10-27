package db

import (
	"errors"

	"github.com/vaibhavchalse99/config"
)

type RoleValue string
type ResourceValue string
type CRUDPermissionValues string

const (
	SuperAdmin RoleValue = "SUPER_ADMIN"
	Admin      RoleValue = "ADMIN"
	EndUser    RoleValue = "END_USER"
)

var Roles map[RoleValue][]string

func LoadRoles() {
	Roles = map[RoleValue][]string{
		SuperAdmin: {
			config.CreateUserAdmin,
			config.CreateUser,
			config.GetUsers,
			config.UpdateUser,
			config.DeleteUser,
			config.CreateBook,
			config.GetBooks,
			config.UpdateBook,
			config.DeleteBook,
			config.CreateProfile,
			config.GetProfile,
			config.UpdateProfile,
			config.CreateBookActiviy,
			config.GetBookActivities,
			config.UpdateBookActiviy,
			config.DeleteBookActivity,
		},
		Admin: {
			config.CreateUser,
			config.GetUsers,
			config.CreateBook,
			config.GetBooks,
			config.UpdateBook,
			config.DeleteBook,
			config.CreateProfile,
			config.GetProfile,
			config.UpdateProfile,
			config.CreateBookActiviy,
			config.GetBookActivities,
			config.UpdateBookActiviy,
			config.DeleteBookActivity,
		},
		EndUser: {
			config.GetProfile,
			config.UpdateProfile,
			config.GetBooks,
			config.GetBookActivities,
		},
	}
}

func (r RoleValue) Validate() error {
	switch r {
	case SuperAdmin, Admin, EndUser:
		return nil
	}
	return errors.New("invalid role")
}

func IsAuthorized(role RoleValue, permission string) bool {
	permissions := Roles[role]

	for _, value := range permissions {
		if value == permission {
			return true
		}
	}
	return false
}
