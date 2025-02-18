package models

type SeatClass string

const (
	Economy    SeatClass = "ECONOMY"
	Business   SeatClass = "BUSINESS"
	FirstClass SeatClass = "FIRST"
)

type Seat struct {
	ID         string
	SeatNumber string
	SeatStatus SeatClass
	IsBooked   bool
	Price      float64
}
