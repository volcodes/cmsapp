package router

import (
	"cms-project/internal/navLinks"
	"cms-project/internal/users"

	"github.com/gorilla/mux"
)

// New creates and configures a new router
func New() *mux.Router {
	r := mux.NewRouter()

	// Register routes
	registerUserRoutes(r)
	registerNavLinksRoutes(r)

	return r
}

func registerUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", users.GetUsersHandler).Methods("GET")
}

func registerNavLinksRoutes(r *mux.Router) {
	r.HandleFunc("/nav_links", navLinks.GetNavLinksHandler).Methods("GET")
}

// func registerMenuRoutes(r *mux.Router) {
// 	r.HandleFunc("/menus", menu.GetMenusHandler).Methods("GET")
// 	r.HandleFunc("/menus", menu.CreateMenuHandler).Methods("POST")
// 	r.HandleFunc("/menus/{id}", menu.GetMenuByIDHandler).Methods("GET")
// 	r.HandleFunc("/menus/{id}", menu.DeleteMenuHandler).Methods("DELETE")
// }
