package router

import (
	"github.com/gorilla/mux"
	"github.com/vaibhavchalse99/controller"
)

func Router() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/", controller.SampleFunc)
	r.HandleFunc("/users", controller.AddUser).Methods("POST")
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/login", controller.LoginAPI).Methods("POST")
	r.HandleFunc("/profile/details", controller.IsLoggedIn(controller.GetProfileInfo)).Methods("GET")
	r.HandleFunc("/profile/details", controller.IsLoggedIn(controller.UpdateProfileInfo)).Methods("PUT")
	r.HandleFunc("/books", controller.GetBookList).Methods("GET")
	r.HandleFunc("/books", controller.Addbook).Methods("POST")
	r.HandleFunc("/books/assign", controller.AssignBook).Methods("POST")

	return
}
