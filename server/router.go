package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhavchalse99/api"
	"github.com/vaibhavchalse99/books"
	"github.com/vaibhavchalse99/middlewares"
	"github.com/vaibhavchalse99/users"
)

// const (
// 	versionHeader = "Accept"
// )

func initRouter(dep dependencies) (router *mux.Router) {
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	router = mux.NewRouter()

	router.HandleFunc("/ping", pingHandler)

	router.HandleFunc("/users/superadmin", users.CreateSuperAdmin(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", middlewares.IsLoggedIn(users.CreateUser(dep.UserService), dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", middlewares.IsLoggedIn(users.ListAllUsers(dep.UserService), dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/login", users.UserLogin(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users/profile", middlewares.IsLoggedIn(users.GetProfileDetails(dep.UserService), dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/profile", middlewares.IsLoggedIn(users.UdateProfileDetails(dep.UserService), dep.UserService)).Methods(http.MethodPut)

	router.HandleFunc("/books", middlewares.IsLoggedIn(books.Create(dep.BookService), dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/books", middlewares.IsLoggedIn(books.List(dep.BookService), dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", middlewares.IsLoggedIn(books.GetBookById(dep.BookService), dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/books", middlewares.IsLoggedIn(books.UpdateBook(dep.BookService), dep.UserService)).Methods(http.MethodPut)

	router.HandleFunc("/books/copies/add", middlewares.IsLoggedIn(books.AddBookCopy(dep.BookService), dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/books/copies/remove", middlewares.IsLoggedIn(books.RemoveBookCopy(dep.BookService), dep.UserService)).Methods(http.MethodPost)

	router.HandleFunc("/books/copies/assign", middlewares.IsLoggedIn(books.AssignBook(dep.BookService), dep.UserService)).Methods(http.MethodPost)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
