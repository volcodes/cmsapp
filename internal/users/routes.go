package users

import "github.com/gorilla/mux"

func RegisterUsersRoutes(r *mux.Router) {
	r.HandleFunc("", GetUsersHandler).Methods("GET")
	r.HandleFunc("", CreateUserHandler).Methods("POST")
}
