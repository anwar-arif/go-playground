package models

type UserRole string

const (
	RolePassenger UserRole = "PASSENGER"
	RoleStaff     UserRole = "STAFF"
	RoleAdmin     UserRole = "ADMIN"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     UserRole
	Phone    string
	Address  string
}

type Passenger struct {
	User
	Bookings []Booking
}
