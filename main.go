package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mogueno/burrito-shop/api/handlers"
	auth "github.com/mogueno/burrito-shop/api/middleware"
	database "github.com/mogueno/burrito-shop/utils"
)

func main() {
	r := mux.NewRouter()
	database.InitDB()
	// Register middleware
	r.Use(auth.AuthMiddleware)

	// Register API routes
	r.HandleFunc("/api/burrito", handlers.GetBurritosHandler).Methods("GET")
	r.HandleFunc("/api/orders", handlers.GetOrdersHandler).Methods("GET")
	r.Handle("/api/orders", http.HandlerFunc(handlers.CreateOrderHandler)).Methods("POST")
	r.HandleFunc("/api/orders/{id}", handlers.GetOrderHandler).Methods("GET")
	
	log.Default().Print("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
