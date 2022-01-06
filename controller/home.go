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

func GetHomeScreen(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var home []sm.HomeScreen
	connection.Preload("CurrencyRates").Preload("Links").Find(&home)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(home)
}

func CreateHomeScreen(w http.ResponseWriter, r *http.Request) {
	var home sm.HomeScreen
	err := json.NewDecoder(r.Body).Decode(&home)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateHomeScreen(home)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(home)
}

func EditHomeScreen(w http.ResponseWriter, r *http.Request) {
	var home sm.HomeScreen
	err := json.NewDecoder(r.Body).Decode(&home)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.EditHomeScreen(home)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(home)
}

func DeleteHomeScreen(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteHomeScreen(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
