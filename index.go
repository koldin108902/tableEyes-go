package main

import (
	"net/http"

	"github.com/gorilla/mux"

	router "web/router"
)

func main() {
	app := mux.NewRouter()

	app.HandleFunc("/store", router.GetStore).Methods("GET")
	app.HandleFunc("/store", router.AddStore).Methods("POST")

	http.ListenAndServe(":3100", app)
}
