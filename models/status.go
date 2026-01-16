package models

type SeatStatus string

// enums of seat status
const (
	AVAILABLE SeatStatus = "available"
	BOOKED    SeatStatus = "booked"
	LOCKED    SeatStatus = "locked"
)
