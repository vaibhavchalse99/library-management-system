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
		if err == ErrUserNotExist {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
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
		response, err := service.Login(r.Context(), reqBody)

		if err == ErrUserNotExist {
			api.Error(rw, http.StatusNotFound, api.Response{Message: ErrUserNotExist.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, response)
	})
}

func GetProfileDetails(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := context.Get(r, "user").(User)
		var response response
		response.User = user
		api.Success(rw, http.StatusOK, response)
	})
}

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyEmail || err == errEmptyPassword || err == errEmptyRole
}
