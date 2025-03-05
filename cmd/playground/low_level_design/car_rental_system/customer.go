package car_rental_system

type Customer struct {
	Name             string
	PhoneNo          string
	Email            string
	DrivingLicenseNo string
}

func NewCustomer(name, email, license string) *Customer {
	return &Customer{
		Name:             name,
		PhoneNo:          "",
		Email:            email,
		DrivingLicenseNo: license,
	}
}
