package books

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhavchalse99/api"
)

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var reqBody createBookRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		book, err := service.Create(r.Context(), reqBody)

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, book)
	})
}

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		books, err := service.List(r.Context())
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, books)
	})
}

func GetBookById(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookId := params["bookId"]
		book, err := service.GetBookById(r.Context(), bookId)
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, book)
	})
}
