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

func GetContactScreen(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var contactScreen []sm.ContactScreen
	connection.Preload("Regions").Preload("Regions.Branches").Find(&contactScreen)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactScreen)
}

func CreateContactScreen(w http.ResponseWriter, r *http.Request) {
	var contactScreen sm.ContactScreen
	err := json.NewDecoder(r.Body).Decode(&contactScreen)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateContactScreen(contactScreen)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactScreen)
}

func EditContactScreen(w http.ResponseWriter, r *http.Request) {
	var contactScreen sm.ContactScreen
	err := json.NewDecoder(r.Body).Decode(&contactScreen)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	contactScreen.Id = uint(id)
	database.EditContactScreen(contactScreen)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactScreen)
}

func DeleteContactScreen(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteContactScreen(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
