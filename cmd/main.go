package main

import (
	"cms-project/internal/database"
	"cms-project/router"
	"log"
	"net/http"
)

func main() {
	// Initialize database
	database.InitDB()

	// Initialize router
	r := router.New()

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
