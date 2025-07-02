package categories

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Category struct
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description,omitempty"`
	Image       *string   `json:"image,omitempty"`
	ParentID    *string   `json:"parentId,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// In-memory store for demonstration
var categories = []Category{
	{
		ID:        "1",
		Name:      "Electronics",
		Slug:      "electronics",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// RegisterRoutes registers category routes on the router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/categories", CategoriesHandler).Methods("GET")
	r.HandleFunc("/categories/{id}", CategoryHandler).Methods("GET")
	r.HandleFunc("/categories", CreateCategoryHandler).Methods("POST")
	r.HandleFunc("/categories/{id}", UpdateCategoryHandler).Methods("PUT")
	r.HandleFunc("/categories/{id}", DeleteCategoryHandler).Methods("DELETE")
	// ...other handlers for products and reviews...
}

// Handlers

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, category := range categories {
		if category.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	category.ID = time.Now().Format("20060102150405") // simple unique ID
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	categories = append(categories, category)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, category := range categories {
		if category.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&categories[i]); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			categories[i].UpdatedAt = time.Now()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, category := range categories {
		if category.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
