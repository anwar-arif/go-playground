package car_rental_system

import (
	"sync"
)

type Car struct {
	Make               string
	Model              string
	Year               int
	LicensePlateNumber string
	RentalPricePerDay  float64
	available          bool
	mu                 sync.Mutex
}

func NewCar(make, model string, year int, licensePlateNo string, pricePerDay float64) *Car {
	return &Car{
		Make:               make,
		Model:              model,
		Year:               year,
		LicensePlateNumber: licensePlateNo,
		RentalPricePerDay:  pricePerDay,
		available:          true,
	}
}

func (c *Car) IsAvailable() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.available
}

func (c *Car) SetAvailable(available bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.available = available
}
