package services

import (
	"fmt"
	"itinerary-builder/models"
	"os"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(itinerary models.Itinerary) string {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage()

	// ---------------- Helper Function ----------------
	formatDateTime := func(dt string) string {
		t, err := time.Parse("2006-01-02 15:04", dt)
		if err != nil {
			return dt
		}
		return fmt.Sprintf("%d %s %d\n%s", t.Day(), t.Format("Jan"), t.Year(), t.Format("03:04 PM"))
	}

	// ---------------- Title ----------------
	pdf.SetFont("Arial", "B", 22)
	pdf.CellFormat(0, 14, strings.ToUpper("Itinerary: "+itinerary.Title), "", 1, "C", false, 0, "")
	pdf.Ln(4)

	// ---------------- User & Dates ----------------
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 8, fmt.Sprintf("User: %s", itinerary.UserID), "", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 13)
	pdf.CellFormat(0, 8, fmt.Sprintf("From: %s      To: %s", itinerary.StartDate, itinerary.EndDate), "", 1, "L", false, 0, "")
	pdf.Ln(8)

	// ---------------- Hotels ----------------
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Hotels", "", 1, "L", false, 0, "")
	pdf.Ln(4)

	if len(itinerary.Hotels) > 0 {
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(50, 8, "Hotel Name", "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 8, "City", "1", 0, "C", false, 0, "")
		pdf.CellFormat(35, 8, "Check-in", "1", 0, "C", false, 0, "")
		pdf.CellFormat(35, 8, "Check-out", "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 8, "Nights", "1", 0, "C", false, 0, "")
		pdf.Ln(-1)

		pdf.SetFont("Arial", "", 12)
		for _, h := range itinerary.Hotels {
			city := h.City
			if city == "" {
				city = "-"
			}
			checkIn := h.CheckInDate
			if checkIn == "" {
				checkIn = "-"
			}
			checkOut := h.CheckOutDate
			if checkOut == "" {
				checkOut = "-"
			}

			pdf.CellFormat(50, 8, h.Name, "1", 0, "C", false, 0, "")
			pdf.CellFormat(40, 8, city, "1", 0, "C", false, 0, "")
			pdf.CellFormat(35, 8, checkIn, "1", 0, "C", false, 0, "")
			pdf.CellFormat(35, 8, checkOut, "1", 0, "C", false, 0, "")
			pdf.CellFormat(20, 8, fmt.Sprintf("%d", h.Nights), "1", 0, "C", false, 0, "")
			pdf.Ln(-1)
		}
	} else {
		pdf.SetFont("Arial", "I", 12)
		pdf.CellFormat(0, 8, "No hotel bookings in this itinerary.", "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)

	// ---------------- Flights ----------------
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Flights", "", 1, "L", false, 0, "")
	pdf.Ln(4)

	if len(itinerary.Flights) > 0 {
		colAirline := 28.0
		colFlightNo := 22.0
		colFrom := 30.0
		colTo := 30.0
		colDepart := 32.0
		colArrival := 32.0
		lineHeight := 6.0

		// Header
		pdf.SetFont("Arial", "B", 11)
		pdf.CellFormat(colAirline, lineHeight*2, "Airline", "1", 0, "C", false, 0, "")
		pdf.CellFormat(colFlightNo, lineHeight*2, "Flight No", "1", 0, "C", false, 0, "")
		pdf.CellFormat(colFrom, lineHeight*2, "From", "1", 0, "C", false, 0, "")
		pdf.CellFormat(colTo, lineHeight*2, "To", "1", 0, "C", false, 0, "")
		pdf.CellFormat(colDepart, lineHeight*2, "Departure", "1", 0, "C", false, 0, "")
		pdf.CellFormat(colArrival, lineHeight*2, "Arrival", "1", 0, "C", false, 0, "")
		pdf.Ln(lineHeight * 2)

		pdf.SetFont("Arial", "", 11)
		for _, f := range itinerary.Flights {
			dep := formatDateTime(f.Departure)
			arr := formatDateTime(f.Arrival)

			depLines := strings.Count(dep, "\n") + 1
			arrLines := strings.Count(arr, "\n") + 1
			maxLines := depLines
			if arrLines > maxLines {
				maxLines = arrLines
			}
			rowHeight := float64(maxLines) * lineHeight

			x := pdf.GetX()
			y := pdf.GetY()

			pdf.CellFormat(colAirline, rowHeight, f.Airline, "1", 0, "C", false, 0, "")
			pdf.CellFormat(colFlightNo, rowHeight, f.FlightNo, "1", 0, "C", false, 0, "")
			pdf.CellFormat(colFrom, rowHeight, f.From, "1", 0, "C", false, 0, "")
			pdf.CellFormat(colTo, rowHeight, f.To, "1", 0, "C", false, 0, "")

			pdf.SetXY(x+colAirline+colFlightNo+colFrom+colTo, y)
			pdf.MultiCell(colDepart, lineHeight, dep, "1", "C", false)

			pdf.SetXY(x+colAirline+colFlightNo+colFrom+colTo+colDepart, y)
			pdf.MultiCell(colArrival, lineHeight, arr, "1", "C", false)

			pdf.SetXY(x, y+rowHeight)
		}
	} else {
		pdf.SetFont("Arial", "I", 12)
		pdf.CellFormat(0, 8, "No flights booked in this itinerary.", "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)

	// ---------------- Activities ----------------
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Activities", "", 1, "L", false, 0, "")
	pdf.Ln(4)

	if len(itinerary.Activities) > 0 {
		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(50, 8, "Time of Day", "1", 0, "C", false, 0, "")
		pdf.CellFormat(140, 8, "Description / Location", "1", 0, "C", false, 0, "")
		pdf.Ln(-1)

		pdf.SetFont("Arial", "", 12)
		for _, a := range itinerary.Activities {
			timeOfDay := a.TimeOfDay
			if timeOfDay == "" {
				timeOfDay = "-"
			}
			desc := fmt.Sprintf("%s | %s", a.Description, a.Location)
			pdf.CellFormat(50, 8, strings.Title(timeOfDay), "1", 0, "C", false, 0, "")
			pdf.CellFormat(140, 8, desc, "1", 0, "C", false, 0, "")
			pdf.Ln(-1)
		}
	} else {
		pdf.SetFont("Arial", "I", 12)
		pdf.CellFormat(0, 8, "No activities scheduled for this itinerary.", "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)

	// ---------------- Transfers ----------------
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Transfers", "", 1, "L", false, 0, "")
	pdf.Ln(4)

	if len(itinerary.Transfers) > 0 {
		pdf.SetFont("Arial", "", 12)
		for i, t := range itinerary.Transfers {
			transferDate := ""
			if i == 0 {
				transferDate = itinerary.StartDate
			} else if i == len(itinerary.Transfers)-1 {
				transferDate = itinerary.EndDate
			}
			formattedDate := transferDate
			if transferDate != "" {
				if d, err := time.Parse("2006-01-02", transferDate); err == nil {
					formattedDate = d.Format("02 Jan 2006")
				}
			}
			pdf.CellFormat(0, 8,
				fmt.Sprintf("%s: %s -> %s at %s, %s", t.Mode, t.Pickup, t.Drop, t.Time, formattedDate),
				"", 1, "L", false, 0, "",
			)
		}
	} else {
		pdf.SetFont("Arial", "I", 12)
		pdf.CellFormat(0, 8, "No transfers booked for this itinerary.", "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)

	// ---------------- Payments ----------------
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Payments", "", 1, "L", false, 0, "")
	pdf.Ln(4)

	if len(itinerary.Payments) > 0 {
		pdf.SetFont("Arial", "", 12)
		for _, p := range itinerary.Payments {
			pdf.CellFormat(0, 8, fmt.Sprintf("Installment %d: Rs.%0.2f | Due: %s", p.Installment, p.Amount, p.DueDate), "", 1, "L", false, 0, "")
		}
	} else {
		pdf.SetFont("Arial", "I", 12)
		pdf.CellFormat(0, 8, "No payments recorded for this itinerary.", "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)

	// ---------------- Inclusions / Exclusions ----------------
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "Inclusions / Exclusions", "", 1, "L", false, 0, "")
	pdf.Ln(4)

	pdf.SetFont("Arial", "", 12)
	if len(itinerary.Inclusions) > 0 {
		pdf.CellFormat(0, 8, "Inclusions: "+strings.Join(itinerary.Inclusions, ", "), "", 1, "L", false, 0, "")
	} else {
		pdf.CellFormat(0, 8, "Inclusions: None", "", 1, "L", false, 0, "")
	}
	if len(itinerary.Exclusions) > 0 {
		pdf.CellFormat(0, 8, "Exclusions: "+strings.Join(itinerary.Exclusions, ", "), "", 1, "L", false, 0, "")
	} else {
		pdf.CellFormat(0, 8, "Exclusions: None", "", 1, "L", false, 0, "")
	}
	pdf.Ln(5)

	// ---------------- Ensure output folder ----------------
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		if err := os.MkdirAll("./output", os.ModePerm); err != nil {
			fmt.Println("Error creating output folder:", err)
			return ""
		}
	}

	// ---------------- Save PDF with timestamp ----------------
	timestamp := time.Now().Format("20060102_150405")
	filePath := fmt.Sprintf("./output/%s_%s.pdf", itinerary.ID, timestamp)
	if err := pdf.OutputFileAndClose(filePath); err != nil {
		fmt.Printf("Error generating PDF for itinerary %s: %v\n", itinerary.ID, err)
		return ""
	}

	return filePath
}

