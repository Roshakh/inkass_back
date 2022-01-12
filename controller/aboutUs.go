package controller

import (
	"encoding/json"
	"inkass/inkassback/database"
	"inkass/inkassback/models"
	sm "inkass/inkassback/models/site-models"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAboutUs(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var aboutUs sm.AboutUs
	connection.Preload("Reviews").Find(&aboutUs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aboutUs)
}

func CreateAboutUs(w http.ResponseWriter, r *http.Request) {
	var aboutUs sm.AboutUs
	err := json.NewDecoder(r.Body).Decode(&aboutUs)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateAboutUs(aboutUs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aboutUs)
}

func EditAboutUs(w http.ResponseWriter, r *http.Request) {
	var aboutUs sm.AboutUs
	err := json.NewDecoder(r.Body).Decode(&aboutUs)
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
	aboutUs.Id = uint(id)
	database.EditAboutUs(aboutUs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aboutUs)
}

func DeleteAboutUs(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteAboutUs(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
