package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
	"github.com/vaibhavchalse99/api"
)

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var reqBody createRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
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

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
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
		var reqBody userCredentials
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
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
		user := context.Get(r, "user").(User)
		api.Success(rw, http.StatusOK, user)
	})
}

func UdateProfileDetails(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		var reqBody updateRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		user := context.Get(r, "user").(User)
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
