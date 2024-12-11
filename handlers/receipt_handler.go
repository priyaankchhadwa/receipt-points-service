package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-points-service/models"
	"receipt-points-service/points"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receiptStore = make(map[string]*models.Receipt)
var pointsStore = make(map[string]int)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a unique ID
	id := uuid.New().String()

	// Store the receipt
	receiptStore[id] = &receipt

	// Calculate points
	points := points.CalculatePoints(receipt)
	pointsStore[id] = points

	response := models.ReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := pointsStore[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
