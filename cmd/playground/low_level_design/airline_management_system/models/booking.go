package models

import "time"

type BookingStatus string

const (
	BookingConfirmed BookingStatus = "CONFIRMED"
	BookingCancelled BookingStatus = "CANCELLED"
	BookingRefunded  BookingStatus = "REFUNDED"
)

type PaymentStatus string

const (
	PaymentComplete   PaymentStatus = "COMPLETE"
	PaymentFailed     PaymentStatus = "FAILED"
	PaymentProcessing PaymentStatus = "PROCESSING"
)

type Booking struct {
	ID            string
	Flight        Flight
	Passenger     Passenger
	Seats         []Seat
	Status        BookingStatus
	TotalAmount   float64
	PaymentStatus PaymentStatus
	BaggageInfo   []Baggage
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
