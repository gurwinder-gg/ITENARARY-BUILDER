package handlers

import (
	"encoding/json"
	"fmt"
	"itinerary-builder/models"
	"itinerary-builder/services"
	"net/http"

	"github.com/gorilla/mux"
)

// In-memory storage for itineraries
var itineraries = make(map[string]models.Itinerary)

// CreateItinerary handles POST /itineraries
func CreateItinerary(w http.ResponseWriter, r *http.Request) {
	var itinerary models.Itinerary
	if err := json.NewDecoder(r.Body).Decode(&itinerary); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if itinerary.ID == "" || itinerary.UserID == "" || itinerary.Title == "" {
		http.Error(w, "ID, UserID, and Title are required", http.StatusBadRequest)
		return
	}

	itineraries[itinerary.ID] = itinerary
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(itinerary)
}

// GetItinerary handles GET /itineraries/{id}
func GetItinerary(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	itinerary, exists := itineraries[id]
	if !exists {
		http.Error(w, "Itinerary not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itinerary)
}

// UpdateItinerary handles PUT /itineraries/{id}
func UpdateItinerary(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, exists := itineraries[id]
	if !exists {
		http.Error(w, "Itinerary not found", http.StatusNotFound)
		return
	}

	var updated models.Itinerary
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updated.ID == "" {
		updated.ID = id
	}

	itineraries[id] = updated
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

// DeleteItinerary handles DELETE /itineraries/{id}
func DeleteItinerary(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, exists := itineraries[id]
	if !exists {
		http.Error(w, "Itinerary not found", http.StatusNotFound)
		return
	}

	delete(itineraries, id)
	w.WriteHeader(http.StatusNoContent)
}

// GeneratePDF handles GET /itineraries/{id}/pdf
func GeneratePDF(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	itinerary, exists := itineraries[id]
	if !exists {
		http.Error(w, "Itinerary not found", http.StatusNotFound)
		return
	}

	filePath := services.GeneratePDF(itinerary)
	if filePath == "" {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	// Return the path of the generated PDF
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"pdf_path": filePath})

	fmt.Println("PDF generated at:", filePath)
}
