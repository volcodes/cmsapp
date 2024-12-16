package users

import (
	"encoding/json"
	"net/http"
)

// GetBlogsHandler handles retrieving all blogs
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := GetUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// CreateBlogHandler handles creating a new blog
// func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
// 	var blog Blog
// 	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}
// 	if err := CreateBlog(blog); err != nil {
// 		http.Error(w, "Failed to create blog", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

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
