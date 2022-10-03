package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DiegoLopez-ing/api-gorm-mux/db"
	"github.com/DiegoLopez-ing/api-gorm-mux/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	findUser := db.DB.Find(&users)
	err := findUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var users models.User
	params := mux.Vars(r)
	findByUser := db.DB.First(&users, params["id"])
	err := findByUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	if users.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&users)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Firstname: "Diego",
		LastName:  "Lopez",
		Email:     "diegolopez@mail.com",
	}

	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var users models.User
	params := mux.Vars(r)
	findByUser := db.DB.First(&users, params["id"])
	err := findByUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	if users.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	deleteUser := db.DB.Delete(&users, params["id"])
	errr := deleteUser.Error
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errr.Error()))
	}
	w.WriteHeader(http.StatusOK)
}
