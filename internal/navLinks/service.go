package navLinks

import (
	"cms-project/internal/database"
	"log"
)

// GetNavLinks retrieves all menus from the database
func GetNavLinks() ([]NavLink, error) {
	var navLinks []NavLink
	query := "SELECT * FROM nav_links"
	err := database.DB.Select(&navLinks, query)
	if err != nil {
		log.Printf("Error fetching navLinks: %v", err)
		return nil, err
	}
	return navLinks, nil
}

// // CreateMenu inserts a new menu into the database
// func CreateMenu(NavLink NavLink) error {
// 	query := "INSERT INTO menus (name) VALUES ($1, $2)"
// 	_, err := database.DB.Exec(query, NavLink.Name)
// 	if err != nil {
// 		log.Printf("Error creating menu: %v", err)
// 		return err
// 	}
// 	return nil
// }

// // GetMenuByID retrieves a single menu by its ID
// func GetMenuByID(id int) (*Menu, error) {
// 	var menu Menu
// 	query := "SELECT * FROM menus WHERE id = $1"
// 	err := database.DB.Get(&menu, query, id)
// 	if err != nil {
// 		log.Printf("Error fetching menu by ID: %v", err)
// 		return nil, err
// 	}
// 	return &menu, nil
// }

// // DeleteMenu removes a menu from the database
// func DeleteMenu(id int) error {
// 	query := "DELETE FROM menus WHERE id = $1"
// 	_, err := database.DB.Exec(query, id)
// 	if err != nil {
// 		log.Printf("Error deleting menu: %v", err)
// 		return err
// 	}
// 	return nil
// }
