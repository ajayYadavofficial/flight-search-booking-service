package cerrors

type CError struct {
	Code    string
	Message string
}

func NewCError(code, message string) *CError {
	return &CError{
		Code:    code,
		Message: message,
	}
}

func (e *CError) Error() string {
	return e.Code + ": " + e.Message
}

// ===================== FLIGHT SEARCH SERVICE ERRORS =====================
var (
	// Repository errors
	ErrFlightSearchRepoFailed = NewCError("FLIGHT_SEARCH_REPO_ERROR", "Failed to retrieve flights from repository")
	ErrFlightAddRepoFailed    = NewCError("FLIGHT_ADD_REPO_ERROR", "Failed to add flight to repository")

	// Date format errors
	ErrInvalidDepartureDate = NewCError("INVALID_DEPARTURE_DATE", "Invalid departure date format. Expected: YYYY-MM-DD HH:MM")
	ErrInvalidArrivalDate   = NewCError("INVALID_ARRIVAL_DATE", "Invalid arrival date format. Expected: YYYY-MM-DD HH:MM")
	ErrInvalidSearchDate    = NewCError("INVALID_SEARCH_DATE", "Invalid search date format. Expected: YYYY-MM-DD")

	// Seat initialization error
	ErrFlightSeatInitFailed = NewCError("FLIGHT_SEAT_INIT_ERROR", "Failed to initialize seats for flight")
	ErrNoFlightsFound       = NewCError("NO_FLIGHTS_FOUND", "No flights found matching the search criteria")
)

// ===================== SEAT CONTROLLER SERVICE ERRORS =====================
var (
	// Repository errors
	ErrSeatRepoNotFound     = NewCError("SEAT_REPO_NOT_FOUND", "Flight seats not found in repository")
	ErrSeatRepoAddFailed    = NewCError("SEAT_REPO_ADD_ERROR", "Failed to add seats to repository")
	ErrSeatRepoUpdateFailed = NewCError("SEAT_REPO_UPDATE_ERROR", "Failed to update seats in repository")
	ErrSeatRepoGetFailed    = NewCError("SEAT_REPO_GET_ERROR", "Failed to retrieve seats from repository")

	// Business logic errors
	ErrInsufficientSeats = NewCError("INSUFFICIENT_SEATS", "Not enough available seats for requested quantity")
	ErrSeatLockFailed    = NewCError("SEAT_LOCK_ERROR", "Failed to lock seats")
	ErrSeatUnlockFailed  = NewCError("SEAT_UNLOCK_ERROR", "Failed to unlock seats")
	ErrSeatBookFailed    = NewCError("SEAT_BOOK_ERROR", "Failed to book seats")
)

// ===================== BOOKING SERVICE ERRORS =====================
var (
	// Seat availability errors
	ErrSeatAvailabilityCheckFailed = NewCError("SEAT_AVAILABILITY_ERROR", "Failed to check seat availability")
	ErrSeatNotAvailable            = NewCError("SEAT_NOT_AVAILABLE", "Requested seat(s) are not available")

	// Seat reservation errors
	ErrSeatReservationFailed = NewCError("SEAT_RESERVATION_ERROR", "Failed to reserve seats")

	// Payment errors
	ErrPaymentFailed = NewCError("PAYMENT_FAILURE", "Payment processing failed")

	// Booking repository errors
	ErrBookingRepoFailed = NewCError("BOOKING_REPO_ERROR", "Failed to create booking in repository")
	ErrBookingNotFound   = NewCError("BOOKING_NOT_FOUND", "Booking not found in repository")

	// Seat confirmation errors
	ErrSeatConfirmFailed = NewCError("SEAT_CONFIRM_ERROR", "Failed to confirm seat booking")
)
