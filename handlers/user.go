package handlers

import (
	"encoding/json"
	"net/http"

	"custom-api-server/db"
	"custom-api-server/models"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = uuid.New() // generate a new UUID

	result := db.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	var user models.User
	result := db.DB.First(&user, "id = ?", input.ID)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	db.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID uuid.UUID `json:"id"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	result := db.DB.Delete(&models.User{}, "id = ?", input.ID)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "User not found or could not be deleted", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
