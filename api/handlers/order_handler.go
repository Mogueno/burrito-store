package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/mux"
	"github.com/mogueno/burrito-shop/models"
	"github.com/mogueno/burrito-shop/repository"
)

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders := repository.GetOrders()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the order details
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error decoding order data: %v", err)
		return
	}

	// Validate the order data
	if len(order.Items) == 0 {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	// Save the order
	err = repository.SaveOrder(order)
	if err != nil {
		http.Error(w, "Failed to save the order", http.StatusInternalServerError)
		return
	}

	// Set the response status code and send the response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order created successfully"))
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	order := repository.GetOrder(uint(orderID))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)

}
