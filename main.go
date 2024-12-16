package main

import (
	"cms-project/internal/database"
	"cms-project/router"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize database
	database.InitDB()

	// Initialize router
	r := router.New()

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
