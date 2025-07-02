package products

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Product struct
type Product struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   string    `json:"description"`
	Price         int       `json:"price"`
	OriginalPrice *int      `json:"originalPrice,omitempty"`
	CategoryID    string    `json:"categoryId"`
	Featured      bool      `json:"featured"`
	InStock       bool      `json:"inStock"`
	StockQuantity int       `json:"stockQuantity"`
	SKU           *string   `json:"sku,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// In-memory store for demonstration
var products = []Product{
	{
		ID:            "1",
		Name:          "Sample Product",
		Slug:          "sample-product",
		Description:   "A sample product",
		Price:         100,
		CategoryID:    "1",
		Featured:      false,
		InStock:       true,
		StockQuantity: 10,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
}

// RegisterRoutes registers product routes on the router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", ProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", ProductHandler).Methods("GET")
	r.HandleFunc("/products", CreateProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteProductHandler).Methods("DELETE")
	// ...other handlers for reviews...
}

// Handlers

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	product.ID = time.Now().Format("20060102150405") // simple unique ID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	products = append(products, product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, product := range products {
		if product.ID == id {
			if err := json.NewDecoder(r.Body).Decode(&products[i]); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			products[i].UpdatedAt = time.Now()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
