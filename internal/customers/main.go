package customers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Customer struct
type Customer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// In-memory store for demonstration
var customers = []Customer{
	{
		ID:        "1",
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// RegisterRoutes registers customer routes on the router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/customers", CustomersHandler).Methods("GET")
	r.HandleFunc("/customers/{id}", CustomerHandler).Methods("GET")
	r.HandleFunc("/customers", CreateCustomerHandler).Methods("POST")
	r.HandleFunc("/customers/{id}", UpdateCustomerHandler).Methods("PUT")
	r.HandleFunc("/customers/{id}", DeleteCustomerHandler).Methods("DELETE")
	// ...other handlers for orders and reviews...
}

// Handlers

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, customer := range customers {
		if customer.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	customer.ID = time.Now().Format("20060102150405") // simple unique ID
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	customers = append(customers, customer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, customer := range customers {
		if customer.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&customers[i]); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			customers[i].UpdatedAt = time.Now()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customers[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, customer := range customers {
		if customer.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
