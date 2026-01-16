package flightseatrepo

import (
	"context"
	"flight-search-booking-service/models"
)

// Repository defines the CRUD interface for flight seat data access
// Business logic is handled by the service layer
type Repository interface {
	// GetFlightSeats retrieves flight seat data by flight ID
	GetFlightSeats(ctx context.Context, flightID string) (*models.FlightSeat, error)

	// UpdateFlightSeats updates flight seat data
	UpdateFlightSeats(ctx context.Context, flightID string, flightSeat *models.FlightSeat) error

	// AddFlightSeats adds a new flight with seats to the repository
	AddFlightSeats(ctx context.Context, flightSeat *models.FlightSeat) error

	// RemoveFlightSeats removes a flight from the repository
	RemoveFlightSeats(ctx context.Context, flightID string) error
}
