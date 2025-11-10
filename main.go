package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/carloss765/RestApiGo/db"
	"github.com/carloss765/RestApiGo/routes"
)

func main() {

	db.DBConnection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
