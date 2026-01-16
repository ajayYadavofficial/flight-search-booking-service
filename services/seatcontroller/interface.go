package seatcontroller

import "context"

// SeatController defines the interface for seat management business logic
type SeatController interface {
	AddSeats(ctx context.Context, flightID string, seats int) error
	LockSeats(ctx context.Context, flightID string, seats int) (seatsIds []string, err error)
	UnlockSeats(ctx context.Context, flightID string, seatIds []string) error
	BookSeats(ctx context.Context, flightID string, seats int) error
	GetAvailableSeats(ctx context.Context, flightID string) (int, error)
}
