package main

import (
    "fmt"
    "log"
    "net/http"
    "itinerary-builder/routes"
)

func main() {
    r := routes.RegisterRoutes()
    fmt.Println("🚀 Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
