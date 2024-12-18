package navLinks

import "github.com/gorilla/mux"

func RegisterNavLinksRoutes(r *mux.Router) {
	r.HandleFunc("", GetNavLinksHandler).Methods("GET")
}
