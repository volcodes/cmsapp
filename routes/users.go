package routes

import (
	"net/http"

	"github.com/volcodes/cmsapp/handlers"
)

type UserRouter struct {
	store *handlers.UserStore
}

func NewUserRouter(store *handlers.UserStore) *UserRouter {
	return &UserRouter{
		store: store,
	}
}

func (r *UserRouter) SetupRoutes() {
	http.HandleFunc("/users", r.handleUsers)
	// http.HandleFunc("/item", r.handleItem)
}

func (r *UserRouter) handleUsers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		r.store.GetUsers(w, req)
	case "POST":
		r.store.CreateUser(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// func (r *UserRouter) handleItem(w http.ResponseWriter, req *http.Request) {
// 	switch req.Method {
// 	case "GET":
// 		r.store.GetItem(w, req)
// 	case "PUT":
// 		r.store.UpdateItem(w, req)
// 	case "DELETE":
// 		r.store.DeleteItem(w, req)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }
