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

	// Users routes

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Tasks routes

	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
