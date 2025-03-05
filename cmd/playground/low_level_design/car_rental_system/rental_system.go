package car_rental_system

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type CarRentalSystem struct {
	cars         map[string]*Car
	reservations map[string]*Reservation
	processor    PaymentProcessor
	mu           sync.RWMutex
}

var (
	instance *CarRentalSystem
	once     sync.Once
)

func GetCarRentalSystem() *CarRentalSystem {
	once.Do(func() {
		instance = &CarRentalSystem{
			cars:         make(map[string]*Car),
			reservations: make(map[string]*Reservation),
			processor:    NewCreditCardPaymentProcessor(),
		}
	})
	return instance
}

func (rs *CarRentalSystem) AddCar(car *Car) error {
	if rs.cars[car.LicensePlateNumber] != nil {
		return errors.New("car is already in the system")
	}
	rs.cars[car.LicensePlateNumber] = car
	return nil
}

func (rs *CarRentalSystem) RemoveCar(car *Car) error {
	if rs.cars[car.LicensePlateNumber] == nil {
		return errors.New("car is not in the system")
	}
	delete(rs.cars, car.LicensePlateNumber)
	return nil
}

func (rs *CarRentalSystem) SearchCars(criteria *SearchCriteria) []*Car {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	var results []*Car
	for _, car := range rs.cars {
		if strings.EqualFold(car.Make, criteria.Make) &&
			strings.EqualFold(car.Model, criteria.Model) &&
			car.IsAvailable() &&
			rs.isCarAvailable(car, criteria.StartDate, criteria.EndDate) {
			results = append(results, car)
		}
	}
	return results
}

func (rs *CarRentalSystem) isCarAvailable(car *Car, startDate, endDate time.Time) bool {
	for _, reservation := range rs.reservations {
		if reservation.Car.LicensePlateNumber == car.LicensePlateNumber {
			if !startDate.After(reservation.EndDate) && !endDate.Before(reservation.StartDate) {
				return false
			}
		}
	}
	return true
}

func (rs *CarRentalSystem) MakeReservation(customer *Customer, car *Car, start, end time.Time) (*Reservation, error) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if !rs.isCarAvailable(car, start, end) {
		return nil, fmt.Errorf("car is not available for the selected dates")
	}

	reservationID := rs.generateReservationID()
	reservation := NewReservation(reservationID, customer, car, start, end)
	rs.reservations[reservationID] = reservation
	car.SetAvailable(false)

	return reservation, nil
}

func (rs *CarRentalSystem) CancelReservation(reservationID string) {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	if reservation, exists := rs.reservations[reservationID]; exists {
		reservation.Car.SetAvailable(true)
		delete(rs.reservations, reservationID)
	}
}

func (rs *CarRentalSystem) ProcessPayment(reservation *Reservation) bool {
	return rs.processor.ProcessPayment(reservation.TotalPrice)
}

func (rs *CarRentalSystem) generateReservationID() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return fmt.Sprintf("RES%s", hex.EncodeToString(bytes))
}
