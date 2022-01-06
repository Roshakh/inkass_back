package controller

import (
	"encoding/json"
	"inkass/inkassback/database"
	"inkass/inkassback/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	var reviews []models.Reviews
	connection.Find(&reviews)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func CreateReviews(w http.ResponseWriter, r *http.Request) {
	var reviews models.Reviews
	err := json.NewDecoder(r.Body).Decode(&reviews)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.CreateReviews(reviews)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func EditReviews(w http.ResponseWriter, r *http.Request) {
	var reviews models.Reviews
	err := json.NewDecoder(r.Body).Decode(&reviews)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.EditReviews(reviews)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func DeleteReviews(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		error := models.Error{IsError: true, Message: "Unproccessable entity"}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	database.DeleteReviews(uint(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
