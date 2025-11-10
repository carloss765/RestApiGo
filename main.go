package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/carloss765/RestApiGo/db"
	"github.com/carloss765/RestApiGo/models"
	"github.com/carloss765/RestApiGo/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
