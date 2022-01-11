package controller

import (
	"encoding/json"
	"inkass/inkassback/database"
	"inkass/inkassback/models"
	sm "inkass/inkassback/models/site-models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetService(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var service []sm.Service
	connection.Preload("ServiceCatalog").Preload("SingleService").Find(&service)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func CreateService(w http.ResponseWriter, r *http.Request) {
	var service sm.Service
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateService(service)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func EditService(w http.ResponseWriter, r *http.Request) {
	var service sm.Service
	err := json.NewDecoder(r.Body).Decode(&service)
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
	service.Id = uint(id)
	database.EditService(service)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteService(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
