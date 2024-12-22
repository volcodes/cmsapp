package users

import (
	"cms-project/pkg/response"
	"encoding/json"
	"net/http"
	"strings"
)

// GetBlogsHandler handles retrieving all blogs
// @Summary      Get all users
// @Description  Retrieve a list of all users from the database
// @Tags         User
// @Success      200 {object} response.APIResponse
// @Failure      500 {object} response.APIResponse
// @Router       /api/users [get]
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := GetUsers()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to fetch users", nil)
		return
	}
	response.JSON(w, http.StatusOK, true, "Users retrieved successfully", users)
}

// RegisterHandler handles creating a new user
// @Summary      Creates a new user
// @Description  Create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body User true "User object"
// @Success      201 {object} response.APIResponse
// @Failure      400 {object} response.APIResponse
// @Failure      500 {object} response.APIResponse
// @Router       /api/users/register [post]
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid input format", nil)
		return
	}

	createdUser, err := CreateUser(user)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "required"):
			response.JSON(w, http.StatusBadRequest, false, err.Error(), nil)
		default:
			response.JSON(w, http.StatusInternalServerError, false, "Failed to create user", nil)
		}
		return
	}

	// Remove sensitive data before sending response
	createdUser.Password = ""

	response.JSON(w, http.StatusCreated, true, "User created successfully", createdUser)
}

// LoginHandler handles user login
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        credentials body LoginRequest true "Login credentials"
// @Success      200 {object} response.APIResponse
// @Failure      400 {object} response.APIResponse
// @Failure      401 {object} response.APIResponse
// @Router       /api/users/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid input format", nil)
		return
	}

	user, token, err := Login(loginReq.Email, loginReq.Password)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, false, err.Error(), nil)
		return
	}

	responseData := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	response.JSON(w, http.StatusOK, true, "Login successful", responseData)
}
