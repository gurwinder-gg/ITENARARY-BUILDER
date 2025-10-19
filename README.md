# 🗺️ Itinerary Builder API


A **Golang REST API** for building, managing, and generating day-wise travel itineraries. Supports CRUD operations and PDF generation for travel plans.

---

## ✨ Features

- Full CRUD: Create, Read, Update, Delete itineraries  
- Manage multiple itinerary components:  
  - Hotels (check-in/out, nights)  
  - Flights (departure/arrival, airline, flight number)  
  - Activities (morning, afternoon, evening)  
  - Transfers (pickup/drop, mode, timing)  
- Payments & Inclusions/Exclusions  
- PDF generation stored in `/output` folder  
- User-specific itineraries (`user_id`)  

---

## 🔗 API Endpoints

| Method | Endpoint | Description |
|--------|---------|-------------|
| POST   | `/itineraries` | Create a new itinerary |
| GET    | `/itineraries/{id}` | Retrieve itinerary by ID |
| PUT    | `/itineraries/{id}` | Update itinerary by ID |
| DELETE | `/itineraries/{id}` | Delete itinerary by ID |
| GET    | `/itineraries/{id}/pdf` | Generate and retrieve PDF |

---

## 📦 Sample JSON Request

```json
{
  "id": "trip001",
  "user_id": "gurwinder07",
  "title": "Goa Trip",
  "start_date": "2025-11-01",
  "end_date": "2025-11-05",
  "hotels": [
    {
      "name": "Taj Resort",
      "city": "Goa",
      "check_in_date": "2025-11-01",
      "check_out_date": "2025-11-05",
      "nights": 4
    }
  ],
  "flights": [
    {
      "airline": "IndiGo",
      "flight_no": "6E245",
      "from": "Delhi",
      "to": "Goa",
      "departure": "2025-11-01 06:00",
      "arrival": "2025-11-01 08:30"
    }
  ],
  "activities": [
    {
      "time_of_day": "morning",
      "description": "Beach visit",
      "location": "Baga Beach"
    }
  ],
  "transfers": [
    {
      "mode": "Car",
      "pickup": "Airport",
      "drop": "Hotel",
      "time": "09:00 AM"
    }
  ],
  "payments": [
    {
      "installment": 1,
      "amount": 5000.00,
      "due_date": "2025-10-25"
    }
  ],
  "inclusions": ["Breakfast", "Airport transfer"],
  "exclusions": ["Lunch", "Dinner"]
}

---

Project Structure

itinerary-builder/
├── handlers/       # API request handlers
├── models/         # Data models
├── routes/         # API routes
├── services/       # PDF generation
├── utils/          # Validation helpers
├── output/         # Generated PDFs
├── main.go         # Server entry
├── go.mod
└── go.sum

---

Getting Started

Clone the repo

git clone https://github.com/gurwinder-gg/ITENARARY-BUILDER.git
cd itinerary-builder

---

Install dependencies

go mod tidy

---
Run the server

go run main.go

---
Test APIs
Use Postman or curl to interact with endpoints.
