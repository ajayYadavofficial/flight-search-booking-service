package models

import "time"

// TODO: Ajay
// Segregate Seats from Flight and Create a new Seat Model
// Create Booking Model to manage Bookings

type Airport struct {
	ID               string   `json:"id,omitempty" bson:"_id,omitempty"`
	Code             string   `json:"code,omitempty" bson:"code,omitempty"`
	Name             string   `json:"name,omitempty" bson:"name,omitempty"`
	AdjecentAirports []string `json:"adjecent_airports,omitempty" bson:"adjecent_airports,omitempty"`
	City             string   `json:"city,omitempty" bson:"city,omitempty"`
}

type FlightSearchRequest struct {
	DepartureCode string `json:"departure_code,omitempty" bson:"departure_code,omitempty"`
	ArrivalCode   string `json:"arrival_code,omitempty" bson:"arrival_code,omitempty"`
	// DepartureDate is a string in the format "YYYY-MM-DD"
	DepartureDate    string `json:"departure_date,omitempty" bson:"departure_date,omitempty"`
	IsPriceLowToHigh bool   `json:"is_price_low_to_high,omitempty" bson:"is_price_low_to_high,omitempty"`
	IsDirectFlight   bool   `json:"is_direct_flight,omitempty" bson:"is_direct_flight,omitempty"`
	Passengers       int    `json:"passengers,omitempty" bson:"passengers,omitempty"`
}

type FlightAddRequest struct {
	Airline         string  `json:"airline,omitempty" bson:"airline,omitempty"`
	DepartureCode   string  `json:"departure_code,omitempty" bson:"departure_code,omitempty"`
	ArrivalCode     string  `json:"arrival_code,omitempty" bson:"arrival_code,omitempty"`
	DepartureDate   string  `json:"departure_date,omitempty" bson:"departure_date,omitempty"`
	ArrivalDate     string  `json:"arrival_date,omitempty" bson:"arrival_date,omitempty"`
	TotalSeats      int     `json:"total_seats,omitempty" bson:"total_seats,omitempty"`
	DurationMinutes int     `json:"duration_minutes,omitempty" bson:"duration_minutes,omitempty"`
	Price           float64 `json:"price,omitempty" bson:"price,omitempty"`
}

type Flight struct {
	ID              string    `json:"id,omitempty" bson:"_id,omitempty"`
	Airline         string    `json:"airline,omitempty" bson:"airline,omitempty"`
	DepartureCode   string    `json:"departure_code,omitempty" bson:"departure_code,omitempty"`
	ArrivalCode     string    `json:"arrival_code,omitempty" bson:"arrival_code,omitempty"`
	DepartureTime   time.Time `json:"departure_time,omitempty" bson:"departure_time,omitempty"`
	ArrivalTime     time.Time `json:"arrival_time,omitempty" bson:"arrival_time,omitempty"`
	AvailableSeats  int       `json:"available_seats,omitempty" bson:"available_seats,omitempty"`
	TotalSeats      int       `json:"total_seats,omitempty" bson:"total_seats,omitempty"`
	DurationMinutes int       `json:"duration_minutes,omitempty" bson:"duration_minutes,omitempty"`
	Price           float64   `json:"price,omitempty" bson:"price,omitempty"`
}

type FlightSeat struct {
	FlightID       string `json:"flight_id,omitempty" bson:"_id,omitempty"`
	Seats          []Seat `json:"seats,omitempty" bson:"seats,omitempty"`
	AvailableSeats int    `json:"available_seats,omitempty" bson:"available_seats,omitempty"`
	TotalSeats     int    `json:"total_seats,omitempty" bson:"total_seats,omitempty"`
}

type Seat struct {
	ID     string     `json:"id,omitempty" bson:"_id,omitempty"`
	Status SeatStatus `json:"status,omitempty" bson:"status,omitempty"`
}

type BookingRequest struct {
	UserInfo User   `json:"user_info,omitempty" bson:"user_info,omitempty"`
	FlightID string `json:"flight_id,omitempty" bson:"flight_id,omitempty"`
	Seats    int    `json:"seats,omitempty" bson:"seats,omitempty"`
}

type Booking struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	FlightID    string    `json:"flight_id,omitempty" bson:"flight_id,omitempty"`
	SeatIDs     []string  `json:"seat_ids,omitempty" bson:"seat_ids,omitempty"`
	UserInfo    User      `json:"user_info,omitempty" bson:"user_info,omitempty"`
	SeatsBooked int       `json:"seats_booked,omitempty" bson:"seats_booked,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type User struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Phone string `json:"phone,omitempty" bson:"phone,omitempty"`
}
