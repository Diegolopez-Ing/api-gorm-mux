package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DiegoLopez-ing/api-gorm-mux/db"
	"github.com/DiegoLopez-ing/api-gorm-mux/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Tasks
	findTask := db.DB.Find(&tasks)
	err := findTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	params := mux.Vars(r)
	findByTask := db.DB.First(&task, params["id"])
	err := findByTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	/*{
		Title:       "Hacer la cama",
		Description: "Objetivo diario tender la cama",
		Done:        false,
		UserId:      1,
	}*/

	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var tasks models.Tasks
	params := mux.Vars(r)
	findByTask := db.DB.First(&tasks, params["id"])
	err := findByTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	if tasks.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	deleteTask := db.DB.Delete(&tasks, params["id"])
	errr := deleteTask.Error
	if errr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errr.Error()))
	}
	w.WriteHeader(http.StatusOK)
}
