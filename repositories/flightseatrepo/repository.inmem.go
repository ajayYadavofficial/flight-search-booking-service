package flightseatrepo

import (
	"context"
	"flight-search-booking-service/models"
	"fmt"
)

// TODO: provide Locks
type inMemFlightSeatRepo struct {
	flightSeats map[string]*models.FlightSeat
}

// InMemFlightSeatRepo is an alias for external access in tests
type InMemFlightSeatRepo = inMemFlightSeatRepo

// NewInMemFlightSeatRepo creates a new in-memory flight seat repository
func NewInMemFlightSeatRepo() *inMemFlightSeatRepo {
	return &inMemFlightSeatRepo{
		flightSeats: make(map[string]*models.FlightSeat),
	}
}

// GetFlightSeats retrieves flight seat data by flight ID (CRUD - READ)
func (r *inMemFlightSeatRepo) GetFlightSeats(ctx context.Context, flightID string) (*models.FlightSeat, error) {
	flightSeat, exists := r.flightSeats[flightID]
	if !exists {
		return nil, fmt.Errorf("flight with ID %s not found", flightID)
	}
	return flightSeat, nil
}

// UpdateFlightSeats updates flight seat data (CRUD - UPDATE)
func (r *inMemFlightSeatRepo) UpdateFlightSeats(ctx context.Context, flightID string, flightSeat *models.FlightSeat) error {
	if _, exists := r.flightSeats[flightID]; !exists {
		return fmt.Errorf("flight with ID %s not found", flightID)
	}
	r.flightSeats[flightID] = flightSeat
	return nil
}

// AddFlightSeats adds a flight with seats to the repository (CRUD - CREATE)
func (r *inMemFlightSeatRepo) AddFlightSeats(ctx context.Context, flightSeat *models.FlightSeat) error {
	if flightSeat.FlightID == "" {
		return fmt.Errorf("flight ID cannot be empty")
	}
	r.flightSeats[flightSeat.FlightID] = flightSeat
	return nil
}

// RemoveFlightSeats removes a flight from the repository (CRUD - DELETE)
func (r *inMemFlightSeatRepo) RemoveFlightSeats(ctx context.Context, flightID string) error {
	if _, exists := r.flightSeats[flightID]; !exists {
		return fmt.Errorf("flight with ID %s not found", flightID)
	}
	delete(r.flightSeats, flightID)
	return nil
}
