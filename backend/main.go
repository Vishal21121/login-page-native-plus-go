package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vishal21121/login-page-native-plus-go/controller"
	"github.com/Vishal21121/login-page-native-plus-go/router"
)

func main() {
	// calling the router function from the router package
	route := router.Router()
	fmt.Println("Server is starting")

	// calling the Init() function from the controller package
	controller.Init()

	// listening the server at port 4000
	log.Fatal(http.ListenAndServe(":4000", route))
	fmt.Println("Listening at port 4000")
}
