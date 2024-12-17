package navLinks

import "github.com/gorilla/mux"

func RegisterNavLinksRoutes(r *mux.Router) {
	r.HandleFunc("/api/nav_links", GetNavLinksHandler).Methods("GET")
}
