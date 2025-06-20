package main

import (
	"log"
	"net/http"

	"custom-api-server/db"
)

func main() {
	db.Connect()

	log.Println(" Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
