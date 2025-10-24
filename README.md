# 🗺️ Itinerary Builder API - Assignment for VIGOVIA

A **Golang REST API** for creating, managing, and exporting day-wise travel itineraries.  
Designed for scalability, clarity, and maintainability with clean RESTful endpoints, modular architecture, and dynamic PDF report generation.

---

##  Key Features

**Complete CRUD Operations**  
Easily create, retrieve, update, and delete itineraries.

**Rich Itinerary Components**  
Supports multiple travel entities with structured relationships:
- **Hotels** — check-in/out dates, number of nights  
- **Flights** — airline, flight number, departure & arrival details  
- **Activities** — categorized by morning, afternoon, evening  
- **Transfers** — pickup/drop details, mode, timing  
- **Payments** — amount, currency, status  
- **Inclusions/Exclusions** — manage trip details with precision  

**PDF Export**  
Generate professional day-wise itinerary PDFs using `gofpdf`, stored in the `/output` directory.

**User Scoping**  
Multi-user support with `user_id` association for individual itineraries.

**Mock Persistence Layer**  
In-memory storage simulates database behavior, ideal for prototyping or extending to real DBs.

---

## API Overview

| Method | Endpoint | Description |
|--------|-----------|-------------|
| **POST** | `/itineraries` | Create a new itinerary |
| **GET** | `/itineraries/{id}` | Retrieve itinerary by ID |
| **PUT** | `/itineraries/{id}` | Update itinerary by ID |
| **DELETE** | `/itineraries/{id}` | Delete itinerary by ID |
| **GET** | `/itineraries/{id}/pdf` | Generate & download itinerary PDF |

---

## Tech Stack

| Category | Technology |
|-----------|-------------|
| **Language** | Go (Golang) |
| **Framework / Router** | Gorilla Mux |
| **PDF Generator** | gofpdf |
| **Storage** | In-memory (mock persistence) |
| **Architecture** | RESTful API (modular and extensible design) |

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/gurwinder-gg/ITENARARY-BUILDER.git
cd ITENARARY-BUILDER
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

The API will be available at:  
➡️ **http://localhost:8080**

---

## Project Structure

```
.
├── main.go                 # Entry point of the application
├── handlers/               # API route handlers
├── models/                 # Data models (Itinerary, Flights, Hotels, etc.)
├── pdf/                    # PDF generation logic
├── storage/                # In-memory data storage implementation
├── utils/                  # Helper utilities
└── output/                 # Generated itinerary PDFs
```

---

##  Example Itinerary Object

```json
{
  "id": "trip001",
  "user_id": "user123",
  "title": "Goa Holiday",
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
      "departure": "2025-11-01T06:00",
      "arrival": "2025-11-01T08:30"
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
      "amount": 5000,
      "due_date": "2025-10-25"
    }
  ],
  "inclusions": ["Breakfast", "Airport transfer"],
  "exclusions": ["Lunch", "Dinner"]
}

```

---

##  PDF Output Example

Each generated PDF includes:
- Day-wise itinerary breakdown  
- Hotel, flight, and activity summaries  
- Transfer details and inclusions/exclusions  
- Clean, travel-agency-ready layout  

All PDFs are saved in:  
📁 `/output/{itinerary_id}.pdf`

---

##  Future Enhancements

- 🔹 Integrate a database (PostgreSQL / MongoDB)  
- 🔹 Implement authentication (JWT-based user auth)  
- 🔹 Add file upload support (images, tickets, documents)  
- 🔹 Enhance PDF with branding and templates  
- 🔹 Dockerize for easy deployment  

---

##  Author

**Gurwinder Singh**



