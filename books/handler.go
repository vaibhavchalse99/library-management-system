package books

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/vaibhavchalse99/api"
	"github.com/vaibhavchalse99/config"
	"github.com/vaibhavchalse99/db"
	"github.com/vaibhavchalse99/users"
)

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// decoding req data
		var reqBody createBookRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		//check permission
		loggedInUser := context.Get(r, "user").(users.User)
		hasPermission := db.IsAuthorized(loggedInUser.Role, config.CreateBook)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}

		//create a book
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
		//check permission
		loggedInUser := context.Get(r, "user").(users.User)
		hasPermission := db.IsAuthorized(loggedInUser.Role, config.GetBooks)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}

		//get book list
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
		//extracting params
		params := mux.Vars(r)
		bookId := params["bookId"]

		//check permission
		loggedInUser := context.Get(r, "user").(users.User)
		hasPermission := db.IsAuthorized(loggedInUser.Role, config.GetBooks)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}
		//get book by id
		book, err := service.GetBookById(r.Context(), bookId)
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, book)
	})
}

func UpdateBook(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//decoding the req data
		var reqBody updateBookRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		//checking permission
		loggedInUser := context.Get(r, "user").(users.User)
		hasPermission := db.IsAuthorized(loggedInUser.Role, config.GetBooks)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}

		//updating book
		book, err := service.UpdateBookById(r.Context(), reqBody)
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, book)
	})
}

func AddBookCopy(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//decoding the req data
		var reqbody createBookCopy
		err := json.NewDecoder(r.Body).Decode(&reqbody)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		//checking permission
		loggedInUser := context.Get(r, "user").(users.User)
		hasPermission := db.IsAuthorized(loggedInUser.Role, config.CreateBook)
		if !hasPermission {
			api.Error(rw, http.StatusForbidden, api.Response{Message: "Access Denied"})
			return
		}
		//add bookcopy
		isbn, err := service.AddBookCopy(r.Context(), reqbody)
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		api.Success(rw, http.StatusOK, ISBNResponse{ISBN: isbn})
	})
}
