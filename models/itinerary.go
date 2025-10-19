package models

type Hotel struct {
	Name         string `json:"name"`
	City         string `json:"city"`
	CheckInDate  string `json:"check_in_date"`
	CheckOutDate string `json:"check_out_date"`
	Nights       int    `json:"nights"`
}

type Flight struct {
	Airline   string `json:"airline"`
	FlightNo  string `json:"flight_no"`
	From      string `json:"from"`
	To        string `json:"to"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
}

type Activity struct {
	TimeOfDay   string `json:"time_of_day"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type Transfer struct {
	Mode   string `json:"mode"`
	Pickup string `json:"pickup"`
	Drop   string `json:"drop"`
	Time   string `json:"time"`
}

type Payment struct {
	Installment int     `json:"installment"`
	Amount      float64 `json:"amount"`
	DueDate     string  `json:"due_date"`
}

type Itinerary struct {
	ID         string     `json:"id"`
	UserID     string     `json:"user_id"`
	Title      string     `json:"title"`
	StartDate  string     `json:"start_date"`
	EndDate    string     `json:"end_date"`
	Hotels     []Hotel    `json:"hotels"`
	Flights    []Flight   `json:"flights"`
	Activities []Activity `json:"activities"`
	Transfers  []Transfer `json:"transfers"`
	Payments   []Payment  `json:"payments"`
	Inclusions []string   `json:"inclusions"`
	Exclusions []string   `json:"exclusions"`
	Insurance  []string   `json:"insurance"`
}
