package users

import (
	"cms-project/internal/database"
	"log"
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

// // CreateBlog inserts a new blog into the database
// func CreateBlog(blog Blog) error {
// 	query := "INSERT INTO blogs (title, content) VALUES ($1, $2)"
// 	_, err := database.DB.Exec(query, blog.Title, blog.Content)
// 	if err != nil {
// 		log.Printf("Error creating blog: %v", err)
// 		return err
// 	}
// 	return nil
// }

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
