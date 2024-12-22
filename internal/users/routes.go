package users

import "github.com/gorilla/mux"

func RegisterUsersRoutes(r *mux.Router) {
	r.HandleFunc("", GetUsersHandler).Methods("GET")
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
}
