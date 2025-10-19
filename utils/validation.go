package utils

import (
    "errors"
    "itinerary-builder/models"
)

func ValidateItinerary(itinerary models.Itinerary) error {
    if itinerary.ID == "" || itinerary.UserID == "" || itinerary.Title == "" {
        return errors.New("missing required itinerary fields")
    }
    return nil
}
