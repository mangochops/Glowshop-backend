package orders

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Order struct {
	ID              string    `json:"id"`
	UserID          *string   `json:"userId,omitempty"`
	OrderNumber     string    `json:"orderNumber"`
	Status          string    `json:"status"`
	Subtotal        int       `json:"subtotal"`
	Tax             int       `json:"tax"`
	Shipping        int       `json:"shipping"`
	Total           int       `json:"total"`
	ShippingAddress string    `json:"shippingAddress"`
	BillingAddress  *string   `json:"billingAddress,omitempty"`
	PaymentMethod   string    `json:"paymentMethod"`
	PaymentStatus   string    `json:"paymentStatus"`
	Notes           *string   `json:"notes,omitempty"`
	TrackingNumber  *string   `json:"trackingNumber,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// In-memory store for demonstration
var orders = []Order{
	{
		ID:              "1",
		OrderNumber:     "ORD-001",
		Status:          "PENDING",
		Subtotal:        100,
		Tax:             10,
		Shipping:        5,
		Total:           115,
		ShippingAddress: "123 Main St",
		PaymentMethod:   "card",
		PaymentStatus:   "PENDING",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	},
}

// RegisterRoutes registers order routes on the router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/orders", OrdersHandler).Methods("GET")
	r.HandleFunc("/orders/{id}", OrderHandler).Methods("GET")
	r.HandleFunc("/orders", CreateOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{id}", UpdateOrderHandler).Methods("PUT")
	r.HandleFunc("/orders/{id}", DeleteOrderHandler).Methods("DELETE")
	// ...other handlers for items and reviews...
}

// Handlers

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, order := range orders {
		if order.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	order.ID = time.Now().Format("20060102150405") // simple unique ID
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	orders = append(orders, order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, order := range orders {
		if order.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&orders[i]); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			orders[i].UpdatedAt = time.Now()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(orders[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, order := range orders {
		if order.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
