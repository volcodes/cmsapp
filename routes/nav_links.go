package routes

import (
	"net/http"

	"github.com/volcodes/cmsapp/handlers"
)

type NavLinkRouter struct {
	store *handlers.NavLinkStore
}

func NewNavLinkRouter(store *handlers.NavLinkStore) *NavLinkRouter {
	return &NavLinkRouter{
		store: store,
	}
}

func (r *NavLinkRouter) SetupRoutes() {
	http.HandleFunc("/nav_links", r.handleItems)
	// http.HandleFunc("/item", r.handleItem)
}

func (r *NavLinkRouter) handleItems(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		r.store.GetNavLinks(w, req)
	case "POST":
		r.store.CreateNavLinks(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// func (r *NavLinkRouter) handleItem(w http.ResponseWriter, req *http.Request) {
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
