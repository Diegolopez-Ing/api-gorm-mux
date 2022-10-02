package main

import (
	"net/http"

	"github.com/DiegoLopez-ing/api-gorm-mux/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8080", r)
}
