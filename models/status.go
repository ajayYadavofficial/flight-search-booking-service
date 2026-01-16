package models

type SeatStatus string

const (
	AVAILABLE SeatStatus = "available"
	BOOKED    SeatStatus = "booked"
	LOCKED    SeatStatus = "locked"
)
