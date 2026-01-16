package flightbookingsvc

import (
	"context"
	"flight-search-booking-service/cerrors"
	"flight-search-booking-service/models"
	"flight-search-booking-service/repositories/flightbookingrepo"
	"flight-search-booking-service/services/paymentsvc"
	"flight-search-booking-service/services/seatcontroller"
	"fmt"
	"time"
)

type Booking interface {
	BookFlight(ctx context.Context, bookingReq models.BookingRequest) (bookingId string, isSuccessful bool, err error)
	GetBookingDetails(ctx context.Context, bookingID string) (*models.Booking, error)
}

type bookingSvc struct {
	paymentSvc     paymentsvc.Payment
	seatController seatcontroller.SeatController
	bookingRepo    flightbookingrepo.Repository
}

// NewFlightBookingService creates a new flight booking service
func NewFlightBookingService(paymentSvc paymentsvc.Payment, seatController seatcontroller.SeatController, bookingRepo flightbookingrepo.Repository) *bookingSvc {
	return &bookingSvc{
		paymentSvc:     paymentSvc,
		seatController: seatController,
		bookingRepo:    bookingRepo,
	}
}

// booking logic ,
// get seat availability from seat controller
// reserve seats
// process payment
// confirm booking and update seat status
func (s *bookingSvc) BookFlight(ctx context.Context, bookingReq models.BookingRequest) (bookingId string, isSuccessful bool, err error) {
	// Check seat availability
	available, err := s.seatController.GetAvailableSeats(ctx, bookingReq.FlightID)
	if err != nil {
		return "", false, fmt.Errorf("%w for flight %s: %v", cerrors.ErrSeatAvailabilityCheckFailed, bookingReq.FlightID, err)
	}
	if available < bookingReq.Seats {
		return "", false, fmt.Errorf("%w. Requested: %d, Available: %d for flight %s", cerrors.ErrSeatNotAvailable, bookingReq.Seats, available, bookingReq.FlightID)
	}

	// Reserve seats
	seatIDs, err := s.seatController.LockSeats(ctx, bookingReq.FlightID, bookingReq.Seats)
	if err != nil {
		return "", false, fmt.Errorf("%w: %d seats for flight %s: %v", cerrors.ErrSeatReservationFailed, bookingReq.Seats, bookingReq.FlightID, err)
	}

	transactionId := getNewTransactionID(bookingReq.UserInfo.Name)
	// Process payment
	paymentSuccessful := s.paymentSvc.ProcessPayment(ctx, transactionId, bookingReq.UserInfo.Name)
	if !paymentSuccessful {
		// Release reserved seats in case of payment failure
		_ = s.seatController.UnlockSeats(ctx, bookingReq.FlightID, seatIDs)
		return "", false, fmt.Errorf("%w for transaction %s (user: %s)", cerrors.ErrPaymentFailed, transactionId, bookingReq.UserInfo.Name)
	}

	// Confirm booking
	bookingId, isSuccessful, err = s.bookingRepo.BookFlight(ctx, bookingReq)
	if err != nil {
		// Release reserved seats in case of booking failure
		_ = s.seatController.UnlockSeats(ctx, bookingReq.FlightID, seatIDs)
		return "", false, fmt.Errorf("%w for flight %s: %v", cerrors.ErrBookingRepoFailed, bookingReq.FlightID, err)
	}
	if !isSuccessful {
		// Release reserved seats in case of booking failure
		_ = s.seatController.UnlockSeats(ctx, bookingReq.FlightID, seatIDs)
		return "", false, fmt.Errorf("%w for flight %s", cerrors.ErrBookingRepoFailed, bookingReq.FlightID)
	}

	// Update seat status to booked
	err = s.seatController.BookSeats(ctx, bookingReq.FlightID, len(seatIDs))
	if err != nil {
		return "", false, fmt.Errorf("%w for flight %s: %v", cerrors.ErrSeatConfirmFailed, bookingReq.FlightID, err)
	}

	return bookingId, true, nil
}

// GetBookingDetails retrieves booking details by booking ID
func (s *bookingSvc) GetBookingDetails(ctx context.Context, bookingID string) (*models.Booking, error) {
	booking, err := s.bookingRepo.GetBookingDetails(ctx, bookingID)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %v", cerrors.ErrBookingNotFound, bookingID, err)
	}
	return booking, nil
}

// getNewTransactionID generates a new transaction ID USERNAME_TIMESTAMP_RANDOM
func getNewTransactionID(userName string) string {
	return fmt.Sprintf("%s_%d_%d", userName, time.Now().Unix(), time.Now().UnixNano()%1000)
}
