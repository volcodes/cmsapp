package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/volcodes/cmsapp/db"
	"github.com/volcodes/cmsapp/handlers"
	"github.com/volcodes/cmsapp/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database connection
	connStr := os.Getenv("DATABASE_URL")
	db.InitDB(connStr)
	defer db.CloseDB()

	// Initialize the item store
	ItemStore := handlers.NewNavLinkStore()
	UserStore := handlers.NewUserStore()

	// Create separate routers for different stores
	navLinkRouter := routes.NewNavLinkRouter(ItemStore)
	userRouter := routes.NewUserRouter(UserStore)

	// Setup routes for both routers
	navLinkRouter.SetupRoutes()
	userRouter.SetupRoutes()

	// Start the server
	log.Println("Starting server on port 8080...")
	log.Println("App is running...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
