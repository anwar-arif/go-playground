package car_rental_system

import (
	"time"
)

type Reservation struct {
	ReservationID string
	Customer      *Customer
	Car           *Car
	StartDate     time.Time
	EndDate       time.Time
	TotalPrice    float64
}

func NewReservation(id string, customer *Customer, car *Car, start, end time.Time) *Reservation {
	res := &Reservation{
		ReservationID: id,
		Customer:      customer,
		Car:           car,
		StartDate:     start,
		EndDate:       end,
	}
	res.TotalPrice = res.calculateTotalPrice()
	return res
}

func (res *Reservation) calculateTotalPrice() float64 {
	days := res.EndDate.Sub(res.StartDate).Hours() / 24
	return res.Car.RentalPricePerDay * days
}
