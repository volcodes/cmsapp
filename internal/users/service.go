package users

import (
	"cms-project/internal/database"
	"errors"
	"log"
	"time"
)

// GetUsers retrieves all blogs from the database
func GetUsers() ([]User, error) {
	var user []User
	query := "SELECT * FROM users"
	err := database.DB.Select(&user, query)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}
	return user, nil
}

// CreateUser inserts a new user into the database
func CreateUser(user User) (*User, error) {
	// Validate required fields
	if err := validateUser(user); err != nil {
		return nil, err
	}

	query := `
		INSERT INTO users (
			name, email, created_at, updated_at, deleted_at, 
			status, role, password, phone, address,
			avatar, is_active, is_deleted, is_verified, 
			is_blocked, is_suspended
		) 
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16
		)
		RETURNING id, created_at`

	// Set default values
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Avatar = ""
	user.IsActive = true
	user.Status = "1"

	// Execute query and get the created user's ID
	err := database.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
		user.Status,
		user.Role,
		user.Password, // !! Task: Should be hashed before storage
		user.Phone,
		user.Address,
		user.Avatar,
		user.IsActive,
		user.IsDeleted,
		user.IsVerified,
		user.IsBlocked,
		user.IsSuspended,
	).Scan(&user.ID, &user.CreatedAt)
	log.Printf("A user has been created.")

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return &user, nil
}

// validateUser checks if required fields are present
func validateUser(user User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	// Add more validation as needed
	return nil
}

// // GetBlogByID retrieves a single blog by its ID
// func GetBlogByID(id int) (*Blog, error) {
// 	var blog Blog
// 	query := "SELECT * FROM blogs WHERE id = $1"
// 	err := database.DB.Get(&blog, query, id)
// 	if err != nil {
// 		log.Printf("Error fetching blog by ID: %v", err)
// 		return nil, err
// 	}
// 	return &blog, nil
// }

// // DeleteBlog removes a blog from the database
// func DeleteBlog(id int) error {
// 	query := "DELETE FROM blogs WHERE id = $1"
// 	_, err := database.DB.Exec(query, id)
// 	if err != nil {
// 		log.Printf("Error deleting blog: %v", err)
// 		return err
// 	}
// 	return nil
// }
