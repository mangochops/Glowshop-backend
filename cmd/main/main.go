package main

import (
	"Glowshop/internal/categories"
	"Glowshop/internal/customers"
	"Glowshop/internal/orders"
	"Glowshop/internal/products"
	"Glowshop/internal/user"
	"log"
	"net/http"

	"Glowshop/internal/db"

	"github.com/gorilla/mux"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.Close()

	r := mux.NewRouter()
	categories.RegisterRoutes(r)
	customers.RegisterRoutes(r)
	orders.RegisterRoutes(r)
	products.RegisterRoutes(r)
	user.RegisterRoutes(r)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", enableCORS(r))
}

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Or specify your frontend URL
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}
