package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is a global variable to access the database pool
var DB *pgxpool.Pool

// InitDB initializes the database connection pool
func InitDB(connStr string) {
	var err error
	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}
	log.Println("Database connection established")
}

// CloseDB closes the database connection pool
func CloseDB() {
	DB.Close()
	log.Println("Database connection closed")
}
