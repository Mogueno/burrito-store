package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mogueno/burrito-shop/repository"
)

func GetBurritosHandler(w http.ResponseWriter, r *http.Request) {
	burritos := repository.GetBurritos()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(burritos)
}
