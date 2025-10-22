# 🗺️ Itinerary Builder API - Assignment for VIGOVIA


A **Golang REST API** for building, managing, and generating day wise travel itineraries. Supports CRUD operations and PDF generation for travel plans.

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

## 🛠️ Tech Stack

- **Language:** Go (Golang)  
- **Router:** Gorilla Mux  
- **PDF Generation:** gofpdf  
- **Data Storage:** In-memory (mock persistence)  

---
Project

| Method | Endpoint                | Description                   |
| ------ | ----------------------- | ----------------------------- |
| POST   | `/itineraries`          | Create a new itinerary        |
| GET    | `/itineraries/{id}`     | Retrieve an itinerary by ID   |
| PUT    | `/itineraries/{id}`     | Update an itinerary           |
| DELETE | `/itineraries/{id}`     | Delete an itinerary           |
| GET    | `/itineraries/{id}/pdf` | Generate PDF for an itinerary |

---

## 🚀 Getting Started

### Clone the repository

```bash

git clone https://github.com/gurwinder-gg/ITENARARY-BUILDER.git
cd itinerary-builder
 ---

### Install dependencies
go mod tidy

### Run the server
go run main.go

Server runs on http://localhost:8080 by default.


