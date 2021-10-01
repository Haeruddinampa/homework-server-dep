package router

import (
	"homework-server-depl/controller"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartRoute() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Starting development server at", port)
	log.Println("Quit the server with CTRL-C.")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", controller.Home).Methods("GET")
	myRouter.HandleFunc("/follower/{username}", controller.Follower).Methods("GET")
	myRouter.HandleFunc("/{userid}/detail", controller.DataUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}
