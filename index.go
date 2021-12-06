package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"web/router"
)

func main() {
	app := mux.NewRouter()

	app.HandleFunc("/store", router.GetStore).Methods("GET")
	app.HandleFunc("/store", router.AddStore).Methods("POST")
	app.HandleFunc("/store", router.EditStore).Methods("PUT")
	app.HandleFunc("/store", router.DeleteStore).Methods("DELETE")

	http.ListenAndServe(":3100", app)
}
