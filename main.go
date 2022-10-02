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

	http.ListenAndServe(":8080", r)
}
