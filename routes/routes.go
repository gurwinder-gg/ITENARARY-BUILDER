package routes

import (
    "itinerary-builder/handlers"
    "github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/itineraries", handlers.CreateItinerary).Methods("POST")
    r.HandleFunc("/itineraries/{id}", handlers.GetItinerary).Methods("GET")
    r.HandleFunc("/itineraries/{id}", handlers.UpdateItinerary).Methods("PUT")
    r.HandleFunc("/itineraries/{id}", handlers.DeleteItinerary).Methods("DELETE")
    r.HandleFunc("/itineraries/{id}/pdf", handlers.GeneratePDF).Methods("GET")

    return r
}
