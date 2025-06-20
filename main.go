package main

import (
	"log"
	"net/http"

	"custom-api-server/db"
	"custom-api-server/handlers"
	"custom-api-server/models"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/update", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/delete", handlers.DeleteUser).Methods("DELETE")

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
