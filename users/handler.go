package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vaibhavchalse99/api"
)

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi")
		var reqBody createRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		err = service.create(r.Context(), reqBody)
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

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyEmail || err == errEmptyPassword || err == errEmptyRole
}
