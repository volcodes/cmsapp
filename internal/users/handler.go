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

// CreateUserHandler handles creating a new user
// @Summary      Creates a new user
// @Description  Create a new user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body User true "User object"
// @Success      201 {object} response.APIResponse
// @Failure      400 {object} response.APIResponse
// @Failure      500 {object} response.APIResponse
// @Router       /api/users [post]
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
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

// GetBlogByIDHandler handles retrieving a single blog by ID
// func GetBlogByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
// 		return
// 	}
// 	blog, err := GetBlogByID(id)
// 	if err != nil {
// 		http.Error(w, "Blog not found", http.StatusNotFound)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(blog)
// }

// DeleteBlogHandler handles deleting a blog by ID
// func DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
// 		return
// 	}
// 	if err := DeleteBlog(id); err != nil {
// 		http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }
