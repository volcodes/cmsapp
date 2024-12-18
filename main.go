package MainHandler

import (
	_ "cms-project/docs"
	"cms-project/internal/database"
	"cms-project/router"
	"log"
	"net/http"
	"os"
)

// @title CMS Project API
// @version 1.0
// @description This is a simple CMS API.
// @host ${VERCEL_URL}
// @BasePath /
// @schemes https

// @contact.name Your Name
// @contact.url https://your-website.com
// @contact.email your-email@example.com
func MainHandler() {
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
