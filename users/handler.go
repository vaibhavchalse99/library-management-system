package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
	"github.com/vaibhavchalse99/api"
	"github.com/vaibhavchalse99/config"
	"github.com/vaibhavchalse99/db"
)

func CreateSuperAdmin(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var reqBody createRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		reqBody.Role = db.SuperAdmin
		err = service.Create(r.Context(), reqBody)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
}

func CreateUser(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		//decoding req data
		var reqBody createRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		//check permission
		user := context.Get(r, "user").(User)
		hasPermission := false
		if reqBody.Role == db.Admin {
			hasPermission = db.IsAuthorized(user.Role, config.CreateUserAdmin)

		}
		if reqBody.Role == db.EndUser {
			hasPermission = db.IsAuthorized(user.Role, config.CreateUser)
		}
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}

		//create a new user
		err = service.Create(r.Context(), reqBody)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
}

func ListAllUsers(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//check permission
		user := context.Get(r, "user").(User)
		hasPermission := db.IsAuthorized(user.Role, config.GetUsers)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}

		//get user list
		response, err := service.List(r.Context())
		if err != nil {
			if err == ErrUserNotExist {
				api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			} else {
				api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			}
			return
		}
		api.Success(rw, http.StatusOK, response)
	})
}

func UserLogin(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//decoding req data
		var reqBody userCredentials
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		//get user token
		token, err := service.Login(r.Context(), reqBody)
		if err != nil {
			if err == ErrUserNotExist {
				api.Error(rw, http.StatusNotFound, api.Response{Message: ErrUserNotExist.Error()})
			} else {
				api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			}
			return
		}
		api.Success(rw, http.StatusOK, token)
	})
}

func GetProfileDetails(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//check permission and return profile data
		user := context.Get(r, "user").(User)
		hasPermission := db.IsAuthorized(user.Role, config.GetProfile)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}
		api.Success(rw, http.StatusOK, user)
	})
}

func UdateProfileDetails(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		//decoding req data
		var reqBody updateRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		//check permission
		user := context.Get(r, "user").(User)
		hasPermission := db.IsAuthorized(user.Role, config.UpdateProfile)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}

		//update profile
		updatedUser, err := service.UpdateById(r.Context(), reqBody, user.ID.String())
		if err != nil {
			if err == errInvalidRequest {
				api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			} else {
				api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			}
			return
		}
		api.Success(rw, http.StatusOK, updatedUser)
	})
}

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyEmail || err == errEmptyPassword || err == errEmptyRole
}
