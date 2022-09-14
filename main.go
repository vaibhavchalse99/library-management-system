package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vaibhavchalse99/router"
)

func main() {
	fmt.Println("server started on 4000")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))

}
