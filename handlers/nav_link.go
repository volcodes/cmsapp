package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/volcodes/cmsapp/db"
)

type NavLink struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
	Href  string `json:"href"`
}

// Store represents the in-memory database
type NavLinkStore struct {
	nav_links []NavLink
}

// NewNavLinkStore initializes a new store
func NewNavLinkStore() *NavLinkStore {
	return &NavLinkStore{
		nav_links: make([]NavLink, 0),
	}
}

// GetNavLinks returns all items
func (s *NavLinkStore) GetNavLinks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, label, href FROM nav_links")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var nav_links []NavLink
	for rows.Next() {
		var nav_link NavLink
		if err := rows.Scan(&nav_link.ID, &nav_link.Label, &nav_link.Href); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		nav_links = append(nav_links, nav_link)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nav_links)
}

// CreateItem adds a new item
func (s *NavLinkStore) CreateNavLinks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newItem NavLink
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow(context.Background(),
		"INSERT INTO nav_links (label, href) VALUES ($1, $2) RETURNING id, label, href",
		newItem.Label, newItem.Href).Scan(&newItem.ID, &newItem.Label, &newItem.Href)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// UpdateItem updates an existing item
func (s *NavLinkStore) UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	for i, item := range s.nav_links {
		if fmt.Sprintf("%d", item.ID) == id {
			s.nav_links[i] = item
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

// DeleteItem removes an item
func (s *NavLinkStore) DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	for i, item := range s.nav_links {
		if fmt.Sprintf("%d", item.ID) == id {
			s.nav_links = append(s.nav_links[:i], s.nav_links[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}
