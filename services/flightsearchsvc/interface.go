package flightsearchsvc

import (
	"context"
	"flight-search-booking-service/models"
)

type FlightController interface {
	SearchFlights(ctx context.Context, req models.FlightSearchRequest) ([]models.Flight, error)
	AddFlight(ctx context.Context, flightAddReq models.FlightAddRequest) (flightID string, err error)
}
