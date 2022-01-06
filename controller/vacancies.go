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

func GetVacancies(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var vacancies []sm.Vacancies
	connection.Preload("CityVacancies").Find(&vacancies)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancies)
}

func CreateVacancies(w http.ResponseWriter, r *http.Request) {
	var vacancies sm.Vacancies
	err := json.NewDecoder(r.Body).Decode(&vacancies)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateVacancies(vacancies)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancies)
}

func EditVacancies(w http.ResponseWriter, r *http.Request) {
	var vacancies sm.Vacancies
	err := json.NewDecoder(r.Body).Decode(&vacancies)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.EditVacancies(vacancies)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancies)
}

func DeleteVacancies(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteVacancies(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
