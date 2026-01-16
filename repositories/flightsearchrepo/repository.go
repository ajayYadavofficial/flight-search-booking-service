package flightsearchrepo

import (
	"context"
	"flight-search-booking-service/models"
)

type Repository interface {
	AddFlight(ctx context.Context, flight models.Flight) (flightId string, err error)
	GetFlights(ctx context.Context, req models.FlightSearchRequest) ([]models.Flight, error)
	GetAllFlights(ctx context.Context) ([]models.Flight, error)
	RemoveFlight(ctx context.Context, flightID string) error
}
