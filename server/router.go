package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vaibhavchalse99/api"
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
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
