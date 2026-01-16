package main

import (
	"encoding/json"
	"flight-search-booking-service/models"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// handleSearchFlight handles GET /searchflight
// Query parameters: departure_code, arrival_code, departure_date, is_price_low_to_high, is_direct_flight, passengers
func (s *APIServer) handleSearchFlight(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	queryParams := r.URL.Query()

	// Parse boolean for is_price_low_to_high (default: true)
	isPriceLowToHigh := true
	if val := queryParams.Get("is_price_low_to_high"); val != "" {
		isPriceLowToHigh, _ = strconv.ParseBool(val)
	}

	// Parse boolean for is_direct_flight (default: false)
	isDirectFlight := false
	if val := queryParams.Get("is_direct_flight"); val != "" {
		isDirectFlight, _ = strconv.ParseBool(val)
	}

	// Parse passengers (default: 1)
	passengers := 1
	if val := queryParams.Get("passengers"); val != "" {
		if p, err := strconv.Atoi(val); err == nil {
			passengers = p
		}
	}

	// Create FlightSearchRequest
	searchReq := models.FlightSearchRequest{
		DepartureCode:    queryParams.Get("departure_code"),
		ArrivalCode:      queryParams.Get("arrival_code"),
		DepartureDate:    queryParams.Get("departure_date"),
		IsPriceLowToHigh: isPriceLowToHigh,
		IsDirectFlight:   isDirectFlight,
		Passengers:       passengers,
	}

	// Validate required fields
	if searchReq.DepartureCode == "" || searchReq.ArrivalCode == "" || searchReq.DepartureDate == "" {
		http.Error(w, "Missing required parameters: departure_code, arrival_code, departure_date", http.StatusBadRequest)
		return
	}

	// Call service
	ctx := r.Context()
	flights, err := s.flightSearchService.SearchFlights(ctx, searchReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to search flights: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"count":   len(flights),
		"flights": flights,
	})
}

// handleAddFlight handles POST /addflight
// Request body: FlightAddRequest JSON
func (s *APIServer) handleAddFlight(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var flightAddReq models.FlightAddRequest
	if err := json.NewDecoder(r.Body).Decode(&flightAddReq); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if flightAddReq.Airline == "" || flightAddReq.DepartureCode == "" || flightAddReq.ArrivalCode == "" {
		http.Error(w, "Missing required fields: airline, departure_code, arrival_code", http.StatusBadRequest)
		return
	}

	// Call service
	ctx := r.Context()
	flightID, err := s.flightSearchService.AddFlight(ctx, flightAddReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add flight: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   true,
		"flight_id": flightID,
		"message":   "Flight added successfully",
	})
}

// handleBookFlight handles POST /bookflight
// Request body: BookingRequest JSON
func (s *APIServer) handleBookFlight(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var bookingReq models.BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&bookingReq); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if bookingReq.FlightID == "" || bookingReq.Seats <= 0 {
		http.Error(w, "Missing required fields: flight_id, seats (must be > 0)", http.StatusBadRequest)
		return
	}

	if bookingReq.UserInfo.Name == "" || bookingReq.UserInfo.Email == "" {
		http.Error(w, "Missing required user info: name, email", http.StatusBadRequest)
		return
	}

	// Call service
	ctx := r.Context()
	bookingID, isSuccessful, err := s.bookingService.BookFlight(ctx, bookingReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to book flight: %v", err), http.StatusInternalServerError)
		return
	}

	if !isSuccessful {
		http.Error(w, "Booking failed: either not enough seats available or payment failed", http.StatusBadRequest)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":    true,
		"booking_id": bookingID,
		"message":    "Flight booked successfully",
	})
}

// handleGetBooking handles GET /getbooking
// Query parameters: booking_id
func (s *APIServer) handleGetBooking(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	bookingID := r.URL.Query().Get("booking_id")

	// Validate required fields
	if bookingID == "" {
		http.Error(w, "Missing required parameter: booking_id", http.StatusBadRequest)
		return
	}

	// Call service
	ctx := r.Context()
	booking, err := s.bookingService.GetBookingDetails(ctx, bookingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve booking: %v", err), http.StatusNotFound)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"booking": booking,
	})
}

// handleHealth handles GET /health
func (s *APIServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	})
}
