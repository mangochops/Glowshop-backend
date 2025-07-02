package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID            string     `json:"id"`
	Name          *string    `json:"name,omitempty"`
	Email         string     `json:"email"`
	EmailVerified *time.Time `json:"emailVerified,omitempty"`
	Password      *string    `json:"password,omitempty"`
	Image         *string    `json:"image,omitempty"`
	Role          string     `json:"role"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}

// In-memory store for demonstration
var users = []User{
	{
		ID:        "1",
		Name:      strPtr("John Doe"),
		Email:     "john@example.com",
		Role:      "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// Helper for pointer to string
func strPtr(s string) *string { return &s }

// RegisterRoutes registers user routes on the router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", UsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", UserHandler).Methods("GET")
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")
	// ...other handlers for orders and reviews...
}

// Handlers

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	user.ID = time.Now().Format("20060102150405") // simple unique ID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, user := range users {
		if user.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&users[i]); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			users[i].UpdatedAt = time.Now()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
