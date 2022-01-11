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

func GetTarifTable(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var tarifTable []sm.TarifTable
	connection.Find(&tarifTable)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarifTable)
}

func CreateTarifTable(w http.ResponseWriter, r *http.Request) {
	var tarifTable sm.TarifTable
	err := json.NewDecoder(r.Body).Decode(&tarifTable)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateTarifTable(tarifTable)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarifTable)
}

func EditTarifTable(w http.ResponseWriter, r *http.Request) {
	var tarifTable sm.TarifTable
	err := json.NewDecoder(r.Body).Decode(&tarifTable)
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
	tarifTable.Id = uint(id)
	database.EditTarifTable(tarifTable)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarifTable)
}

func DeleteTarifTable(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteTarifTable(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
