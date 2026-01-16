package flightsearchsvc

import (
	"context"
	"flight-search-booking-service/cerrors"
	"flight-search-booking-service/models"
	"flight-search-booking-service/repositories/flightsearchrepo"
	"flight-search-booking-service/services/seatcontroller"
	"fmt"
	"sort"
	"time"
)

type FlightController interface {
	SearchFlights(ctx context.Context, req models.FlightSearchRequest) ([]models.Flight, error)
	AddFlight(ctx context.Context, flightAddReq models.FlightAddRequest) (flightID string, err error)
}

type flightSearchSvc struct {
	flightSearchRepo  flightsearchrepo.Repository
	seatControllerSvc seatcontroller.SeatController
}

//TODO: Add the functionality to add the flightSeats when creating a flight

// NewFlightSearchService creates a new flight search service
func NewFlightSearchService(flightSearchRepo flightsearchrepo.Repository, seatControllerSvc seatcontroller.SeatController) *flightSearchSvc {
	return &flightSearchSvc{
		flightSearchRepo:  flightSearchRepo,
		seatControllerSvc: seatControllerSvc,
	}
}

// SearchFlights retrieves flights based on search criteria (direct and connecting flights)
func (s *flightSearchSvc) SearchFlights(ctx context.Context, req models.FlightSearchRequest) ([]models.Flight, error) {
	// Get all flights from repository
	allFlights, err := s.flightSearchRepo.GetAllFlights(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", cerrors.ErrFlightSearchRepoFailed, err)
	}

	var results []models.Flight

	// Get direct flights
	directFlights := s.getDirectFlights(allFlights, req)
	results = append(results, directFlights...)

	// Get connecting flights if not specifically requesting direct flights only
	if !req.IsDirectFlight {
		connectingFlights := s.getConnectingFlights(allFlights, req)
		results = append(results, connectingFlights...)
	}

	// Sort results based on price preference
	if req.IsPriceLowToHigh {
		sort.Slice(results, func(i, j int) bool {
			return results[i].Price < results[j].Price
		})
	} else {
		sort.Slice(results, func(i, j int) bool {
			return results[i].Price > results[j].Price
		})
	}

	// Populate available seats for each flight from seat controller
	for i, flight := range results {
		availableSeats, err := s.seatControllerSvc.GetAvailableSeats(ctx, flight.ID)
		if err != nil {
			// Log error but continue with other flights
			fmt.Printf("Failed to get available seats for flight %s: %v\n", flight.ID, err)
			results[i].AvailableSeats = 0
		} else {
			results[i].AvailableSeats = availableSeats
		}
	}
	if len( results ) == 0 {
		return nil, cerrors.ErrNoFlightsFound
	}

	return results, nil
}

// getDirectFlights finds direct flights matching the search criteria
func (s *flightSearchSvc) getDirectFlights(flights []models.Flight, req models.FlightSearchRequest) []models.Flight {
	var directFlights []models.Flight

	// Parse departure date string to time.Time
	requestDate, err := parseDate(req.DepartureDate)
	if err != nil {
		// If parsing fails, return empty results
		return directFlights
	}

	for _, flight := range flights {
		// Check if flight matches search criteria
		if flight.DepartureCode == req.DepartureCode &&
			flight.ArrivalCode == req.ArrivalCode &&
			isSameDate(flight.DepartureTime, requestDate) &&
			flight.TotalSeats >= req.Passengers {
			directFlights = append(directFlights, flight)
		}
	}

	return directFlights
}

// getConnectingFlights finds one-stop connecting flights
func (s *flightSearchSvc) getConnectingFlights(flights []models.Flight, req models.FlightSearchRequest) []models.Flight {
	var connectingFlights []models.Flight

	// Parse departure date string to time.Time
	requestDate, err := parseDate(req.DepartureDate)
	if err != nil {
		// If parsing fails, return empty results
		return connectingFlights
	}

	// Find all flights from origin on the requested date
	for _, firstFlight := range flights {
		if firstFlight.DepartureCode != req.DepartureCode ||
			!isSameDate(firstFlight.DepartureTime, requestDate) ||
			firstFlight.TotalSeats < req.Passengers {
			continue
		}

		// Find connecting flights from the first flight's destination to the final destination
		for _, secondFlight := range flights {
			// Check if second flight connects from first flight's arrival to final destination
			if secondFlight.DepartureCode == firstFlight.ArrivalCode &&
				secondFlight.ArrivalCode == req.ArrivalCode &&
				isSameDate(secondFlight.DepartureTime, requestDate) &&
				secondFlight.TotalSeats >= req.Passengers {

				// Create a composite flight for the connecting itinerary
				compositeFlight := models.Flight{
					ID:              firstFlight.ID + "-" + secondFlight.ID, // Composite ID
					Airline:         firstFlight.Airline + "/" + secondFlight.Airline,
					DepartureCode:   firstFlight.DepartureCode,
					ArrivalCode:     secondFlight.ArrivalCode,
					DepartureTime:   firstFlight.DepartureTime,
					ArrivalTime:     secondFlight.ArrivalTime,
					TotalSeats:      minSeats(firstFlight.TotalSeats, secondFlight.TotalSeats),
					DurationMinutes: firstFlight.DurationMinutes + secondFlight.DurationMinutes,
					Price:           firstFlight.Price + secondFlight.Price,
				}

				connectingFlights = append(connectingFlights, compositeFlight)
			}
		}
	}

	return connectingFlights
}

// minSeats returns the minimum of two seat counts
func minSeats(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// parseDate parses a date string in format "YYYY-MM-DD" to time.Time
func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// isSameDate checks if a time.Time and a date time.Time are on the same day
func isSameDate(flightTime, requestDate time.Time) bool {
	return flightTime.Year() == requestDate.Year() &&
		flightTime.Month() == requestDate.Month() &&
		flightTime.Day() == requestDate.Day()
}

// AddFlight adds a flight
func (s *flightSearchSvc) AddFlight(ctx context.Context, flightAddReq models.FlightAddRequest) (flightID string, err error) {
	// Parse date strings to time.Time
	departureTime, err := time.Parse("2006-01-02 15:04", flightAddReq.DepartureDate)
	if err != nil {
		return "", err
	}

	arrivalTime, err := time.Parse("2006-01-02 15:04", flightAddReq.ArrivalDate)
	if err != nil {
		return "", err
	}

	// Create Flight object from FlightAddRequest
	flight := models.Flight{
		Airline:         flightAddReq.Airline,
		DepartureCode:   flightAddReq.DepartureCode,
		ArrivalCode:     flightAddReq.ArrivalCode,
		DepartureTime:   departureTime,
		ArrivalTime:     arrivalTime,
		TotalSeats:      flightAddReq.TotalSeats,
		DurationMinutes: flightAddReq.DurationMinutes,
		Price:           flightAddReq.Price,
	}

	// Add flight to repository
	flightID, err = s.flightSearchRepo.AddFlight(ctx, flight)
	if err != nil {
		return "", err
	}

	// Add seats for the flight
	err = s.seatControllerSvc.AddSeats(ctx, flightID, flight.TotalSeats)
	if err != nil {
		return "", fmt.Errorf("failed to add seats for flight %s: %v", flightID, err)
	}

	return flightID, nil
}
