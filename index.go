package main

import (
	"net/http"

	"github.com/gorilla/mux"

	router "web/router"
)

func main() {
	// mux := http.NewServeMux()
	app := mux.NewRouter()

	app.HandleFunc("/store", router.GetStore).Methods("GET")

	http.ListenAndServe(":3100", app)
}
