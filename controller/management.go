package controller

import (
	"encoding/json"
	"inkass/inkassback/database"
	"inkass/inkassback/models"
	ms "inkass/inkassback/models/site-models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetManagement(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var management []ms.Management
	connection.Preload("Biography").Preload("WorkTime").Find(&management)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(management)
}

func CreateManagement(w http.ResponseWriter, r *http.Request) {
	var management ms.Management
	err := json.NewDecoder(r.Body).Decode(&management)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateManagement(management)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(management)
}

func EditManagement(w http.ResponseWriter, r *http.Request) {
	var management ms.Management
	err := json.NewDecoder(r.Body).Decode(&management)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	management.Id = uint(id)
	database.EditManagement(management)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(management)
}

func DeleteManagement(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteManagement(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
