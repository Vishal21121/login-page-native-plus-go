package router

import (
	"github.com/Vishal21121/login-page-native-plus-go/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	// calling the NewRouter function
	router := mux.NewRouter()
	// Handling the /create route and accepting the method POST
	router.HandleFunc("/create", controller.Createuser).Methods("POST")
	return router
}
