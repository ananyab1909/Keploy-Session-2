package main

import (
	"log"
	"net/http"

	"custom-api-server/config"
	"custom-api-server/db"
)

func main() {
	config.LoadEnv()
	db.Connect()

	log.Println("Server is running on port", config.GetEnv("PORT", "8080"))
	http.ListenAndServe(":"+config.GetEnv("PORT", "8080"), nil)
}
