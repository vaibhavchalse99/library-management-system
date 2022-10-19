package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhavchalse99/api"
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

	router.HandleFunc("/users", users.Create(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", users.List(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/login", users.UserLogin(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users/profile", middlewares.IsLoggedIn(users.GetProfileDetails(dep.UserService), dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/profile", middlewares.IsLoggedIn(users.UdateProfileDetails(dep.UserService), dep.UserService)).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
