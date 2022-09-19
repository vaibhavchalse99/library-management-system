package server

import (
	"fmt"
	"strconv"

	"github.com/urfave/negroni"
	"github.com/vaibhavchalse99/config"
)

func StartAPIServer() {
	port := config.AppPort()
	server := negroni.Classic()

	dependencies, err := initDependencies()
	if err != nil {
		panic(err)
	}

	router := initRouter(dependencies)
	server.UseHandler(router)

	server.Run(fmt.Sprintf(":%s", strconv.Itoa(port)))

	// fmt.Println("server started on 4000")
	// r := router.Router()
	// log.Fatal(http.ListenAndServe(":4000", r))
}
