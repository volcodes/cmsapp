package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/volcodes/cmsapp/db"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Store represents the in-memory database
type UserStore struct {
	users []User
}

// NewUserStore initializes a new store
func NewUserStore() *UserStore {
	return &UserStore{
		users: make([]User, 0),
	}
}

// GetUsers returns all users
func (s *UserStore) GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(context.Background(), "SELECT name, email, created_at FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Name, &user.Email, &user.CreatedAt); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser adds a new user
func (s *UserStore) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newItem User
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.DB.QueryRow(context.Background(),
		"INSERT INTO users (name, email, created_at) VALUES ($1, $2) RETURNING id, name, email, created_at",
		newItem.Name, newItem.Email, newItem.CreatedAt).Scan(&newItem.Name, &newItem.Email, &newItem.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// UpdateItem updates an existing item
// func (s *UserStore) UpdateItem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id := r.URL.Query().Get("id")
// 	for i, item := range s.users {
// 		if fmt.Sprintf("%d", item.ID) == id {
// 			s.users[i] = item
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}
// }

// // DeleteItem removes an item
// func (s *UserStore) DeleteItem(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id := r.URL.Query().Get("id")
// 	for i, item := range s.users {
// 		if fmt.Sprintf("%d", item.ID) == id {
// 			s.users = append(s.users[:i], s.users[i+1:]...)
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}
// 	http.Error(w, "Item not found", http.StatusNotFound)
// }
