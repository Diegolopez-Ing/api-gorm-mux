package main

import (
	"net/http"

	"github.com/DiegoLopez-ing/api-gorm-mux/db"
	"github.com/DiegoLopez-ing/api-gorm-mux/models"
	"github.com/DiegoLopez-ing/api-gorm-mux/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConection()
	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.User{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
