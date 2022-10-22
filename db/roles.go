package db

type CRUDPermissions struct {
	Create, Read, Update, Delete bool
}

type Resources struct {
	User, Book, BookActivity, BookReport CRUDPermissions
}

type Roles struct {
	SuperAdmin, Admin, EndUser Resources
}

var roles Roles

func LoadRoles() {
	roles = Roles{
		SuperAdmin: Resources{
			User:         CRUDPermissions{true, true, true, true},
			Book:         CRUDPermissions{true, true, true, true},
			BookActivity: CRUDPermissions{true, true, true, true},
		},
		Admin: Resources{
			User:         CRUDPermissions{true, true, false, false},
			Book:         CRUDPermissions{true, true, true, true},
			BookActivity: CRUDPermissions{true, true, true, true},
		},
		EndUser: Resources{
			User:         CRUDPermissions{false, true, true, false},
			Book:         CRUDPermissions{false, true, false, false},
			BookActivity: CRUDPermissions{false, true, false, false},
		},
	}
}

type RoleValue string
type ResourceValue string

var (
	SuperAdmin RoleValue = "SUPER_ADMIN"
	Admin      RoleValue = "ADMIN"
	EndUser    RoleValue = "END_USER"
)

var (
	UserResource         ResourceValue = "User"
	BookResource         ResourceValue = "Book"
	BookActivityResource ResourceValue = "BookActivity"
)

func GetPermissions(role RoleValue, resource ResourceValue) (permissions CRUDPermissions) {
	if role == SuperAdmin {
		if resource == UserResource {
			permissions = roles.SuperAdmin.User
		}
		if resource == BookResource {
			permissions = roles.SuperAdmin.User
		}
		if resource == BookActivityResource {
			permissions = roles.SuperAdmin.User
		}
	}
	if role == Admin {
		if resource == UserResource {
			permissions = roles.Admin.User
		}
		if resource == BookResource {
			permissions = roles.Admin.User
		}
		if resource == BookActivityResource {
			permissions = roles.Admin.User
		}
	}
	if role == EndUser {
		if resource == UserResource {
			permissions = roles.EndUser.User
		}
		if resource == BookResource {
			permissions = roles.EndUser.User
		}
		if resource == BookActivityResource {
			permissions = roles.EndUser.User
		}
	}
	return
}
