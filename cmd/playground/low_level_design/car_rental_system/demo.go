package car_rental_system

import (
	"fmt"
	"time"
)

func Run() {
	rentalSystem := GetCarRentalSystem()

	// Add cars to the rental system
	rentalSystem.AddCar(NewCar("Toyota", "Camry", 2022, "ABC123", 50.0))
	rentalSystem.AddCar(NewCar("Honda", "Civic", 2021, "XYZ789", 45.0))
	rentalSystem.AddCar(NewCar("Ford", "Mustang", 2023, "DEF456", 80.0))

	// Create customer
	customer := NewCustomer("John Doe", "john@example.com", "DL1234")

	// Make reservation
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 3)

	availableCars := rentalSystem.SearchCars(NewSearchCriteria("Toyota", "Camry", 40, 60, startDate, endDate))
	if len(availableCars) > 0 {
		selectedCar := availableCars[0]
		reservation, err := rentalSystem.MakeReservation(customer, selectedCar, startDate, endDate)
		if err == nil {
			if rentalSystem.ProcessPayment(reservation) {
				fmt.Printf("Reservation successful. Reservation ID: %s\n", reservation.ReservationID)
			} else {
				fmt.Println("Payment failed. Reservation canceled.")
				rentalSystem.CancelReservation(reservation.ReservationID)
			}
		} else {
			fmt.Printf("Reservation failed: %v\n", err)
		}
	} else {
		fmt.Println("No available cars found for the given criteria.")
	}
}
