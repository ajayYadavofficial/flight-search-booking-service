package flightsearchrepo

import (
	"context"
	"flight-search-booking-service/models"
	"fmt"
)

type inMemFlightSearchRepo struct {
	flights map[string]models.Flight
}

// NewInMemFlightSearchRepo creates a new in-memory flight search repository
func NewInMemFlightSearchRepo() *inMemFlightSearchRepo {
	return &inMemFlightSearchRepo{
		flights: make(map[string]models.Flight),
	}
}

// AddFlight adds a flight to the repository
func (r *inMemFlightSearchRepo) AddFlight(ctx context.Context, flight models.Flight) error {
	if flight.ID == "" {
		return fmt.Errorf("flight ID cannot be empty")
	}
	r.flights[flight.ID] = flight
	return nil
}

// GetFlights returns all flights (filtering logic moved to service layer)
func (r *inMemFlightSearchRepo) GetFlights(ctx context.Context, req models.FlightSearchRequest) ([]models.Flight, error) {
	return r.GetAllFlights(ctx)
}

// GetAllFlights returns all flights in the repository
func (r *inMemFlightSearchRepo) GetAllFlights(ctx context.Context) ([]models.Flight, error) {
	var flights []models.Flight
	for _, flight := range r.flights {
		flights = append(flights, flight)
	}
	return flights, nil
}

// RemoveFlight removes a flight from the repository by ID
func (r *inMemFlightSearchRepo) RemoveFlight(ctx context.Context, flightID string) error {
	if _, exists := r.flights[flightID]; !exists {
		return fmt.Errorf("flight with ID %s not found", flightID)
	}
	delete(r.flights, flightID)
	return nil
}
