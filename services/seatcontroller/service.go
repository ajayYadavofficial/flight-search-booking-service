package seatcontroller

import (
	"context"
	"flight-search-booking-service/cerrors"
	"flight-search-booking-service/models"
	"flight-search-booking-service/repositories/flightseatrepo"
	"fmt"
)

// seatcontrollersvc implements the SeatController interface
type seatcontrollersvc struct {
	repo flightseatrepo.Repository
}

// NewSeatControllerService creates a new seat controller service
func NewSeatControllerService(repo flightseatrepo.Repository) SeatController {
	return &seatcontrollersvc{
		repo: repo,
	}
}

// AddSeats adds seats to a flight
// Business Logic:
// 1. Create seat entries
// 2. Initialize available seat count
// 3. Save to repository
func (s *seatcontrollersvc) AddSeats(ctx context.Context, flightID string, seats int) error {
	seatList := make([]models.Seat, 0, seats)
	for i := 0; i < seats; i++ {
		seatList = append(seatList, models.Seat{
			ID:     fmt.Sprintf("%d", i+1),
			Status: models.AVAILABLE,
		})
	}

	flightSeat := &models.FlightSeat{
		FlightID:       flightID,
		Seats:          seatList,
		AvailableSeats: seats,
	}

	err := s.repo.AddFlightSeats(ctx, flightSeat)
	if err != nil {
		return fmt.Errorf("%w for flight %s: %v", cerrors.ErrSeatRepoAddFailed, flightID, err)
	}

	return nil
}

// LockSeats locks a number of seats for a flight
// Business Logic:
// 1. Validate flight exists
// 2. Check if enough available seats
// 3. Find available seats and mark as locked
// 4. Update available seat count
// 5. Return locked seat IDs
func (s *seatcontrollersvc) LockSeats(ctx context.Context, flightID string, seats int) (seatsIds []string, err error) {
	// Fetch the flight seat data from repository
	flightSeat, err := s.repo.GetFlightSeats(ctx, flightID)
	if err != nil {
		return nil, fmt.Errorf("%w for flight %s: %v", cerrors.ErrSeatRepoNotFound, flightID, err)
	}

	// Validate seat count
	if flightSeat.AvailableSeats < seats {
		return nil, fmt.Errorf("%w. Requested: %d, Available: %d for flight %s", cerrors.ErrInsufficientSeats, seats, flightSeat.AvailableSeats, flightID)
	}

	// Business Logic: Find available seats and lock them
	lockedSeats := 0
	for i := range flightSeat.Seats {
		if flightSeat.Seats[i].Status != models.LOCKED && flightSeat.Seats[i].Status != models.BOOKED {
			flightSeat.Seats[i].Status = models.LOCKED
			seatsIds = append(seatsIds, flightSeat.Seats[i].ID)
			lockedSeats++
			if lockedSeats == seats {
				break
			}
		}
	}

	// Update available seat count
	flightSeat.AvailableSeats -= lockedSeats

	// Save the updated flight seat data back to repository
	err = s.repo.UpdateFlightSeats(ctx, flightID, flightSeat)
	if err != nil {
		return nil, fmt.Errorf("%w after locking for flight %s: %v", cerrors.ErrSeatRepoUpdateFailed, flightID, err)
	}

	return seatsIds, nil
}

// UnlockSeats unlocks previously locked seats for a flight
// Business Logic:
// 1. Validate flight exists
// 2. Find locked seats matching the provided seat IDs
// 3. Mark them as available
// 4. Update available seat count
func (s *seatcontrollersvc) UnlockSeats(ctx context.Context, flightID string, seatIds []string) error {
	// Fetch the flight seat data from repository
	flightSeat, err := s.repo.GetFlightSeats(ctx, flightID)
	if err != nil {
		return fmt.Errorf("%w for flight %s: %v", cerrors.ErrSeatRepoNotFound, flightID, err)
	}

	// Business Logic: Find locked seats and unlock them
	unlockedSeats := 0
	for _, seatID := range seatIds {
		for i := range flightSeat.Seats {
			if flightSeat.Seats[i].ID == seatID && flightSeat.Seats[i].Status == models.LOCKED {
				flightSeat.Seats[i].Status = models.AVAILABLE
				unlockedSeats++
				break
			}
		}
	}

	// Update available seat count
	flightSeat.AvailableSeats += unlockedSeats

	// Save the updated flight seat data back to repository
	err = s.repo.UpdateFlightSeats(ctx, flightID, flightSeat)
	if err != nil {
		return fmt.Errorf("%w after unlocking for flight %s: %v", cerrors.ErrSeatRepoUpdateFailed, flightID, err)
	}

	return nil
}

// BookSeats books (confirms) a number of locked seats for a flight
// Business Logic:
// 1. Validate flight exists
// 2. Find locked seats (that were previously locked)
// 3. Mark them as booked
// 4. Note: Available seat count is NOT changed (was already reduced during lock)
func (s *seatcontrollersvc) BookSeats(ctx context.Context, flightID string, seats int) error {
	// Fetch the flight seat data from repository
	flightSeat, err := s.repo.GetFlightSeats(ctx, flightID)
	if err != nil {
		return fmt.Errorf("%w for flight %s: %v", cerrors.ErrSeatRepoNotFound, flightID, err)
	}

	// Business Logic: Find locked seats and mark as booked
	bookedSeats := 0
	for i := range flightSeat.Seats {
		if flightSeat.Seats[i].Status == models.LOCKED {
			flightSeat.Seats[i].Status = models.BOOKED
			bookedSeats++
			if bookedSeats == seats {
				break
			}
		}
	}

	// Save the updated flight seat data back to repository
	err = s.repo.UpdateFlightSeats(ctx, flightID, flightSeat)
	if err != nil {
		return fmt.Errorf("%w after booking for flight %s: %v", cerrors.ErrSeatRepoUpdateFailed, flightID, err)
	}

	return nil
}

// GetAvailableSeats returns the number of available seats for a flight
// Business Logic:
// 1. Fetch flight seat data from repository
// 2. Return the available seat count
func (s *seatcontrollersvc) GetAvailableSeats(ctx context.Context, flightID string) (int, error) {
	flightSeat, err := s.repo.GetFlightSeats(ctx, flightID)
	if err != nil {
		return 0, fmt.Errorf("%w for flight %s: %v", cerrors.ErrSeatRepoGetFailed, flightID, err)
	}

	return flightSeat.AvailableSeats, nil
}
