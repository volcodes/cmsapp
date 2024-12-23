package users

import (
	"cms-project/internal/database"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func Login(email, password string) (*User, string, error) {
	var user User
	query := "SELECT * FROM users WHERE email = $1 AND is_active = true"

	err := database.DB.Get(&user, query, email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// TODO: Replace with proper password hashing comparison
	if user.Password != password {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with your secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	// Clear sensitive data
	user.Password = ""
	log.Printf("Logged %v, token: %v", user.Email, tokenString)

	return &user, tokenString, nil
}
