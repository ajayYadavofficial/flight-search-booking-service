package flightbookingsvc

import (
	"context"
	"flight-search-booking-service/models"
)

type Booking interface {
	BookFlight(ctx context.Context, bookingReq models.BookingRequest) (bookingId string, isSuccessful bool, err error)
	GetBookingDetails(ctx context.Context, bookingID string) (*models.Booking, error)
}
