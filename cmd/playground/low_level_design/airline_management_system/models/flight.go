package models

import (
	"sync"
	"time"
)

type FlightStatus string

const (
	Scheduled FlightStatus = "SCHEDULED"
	Delayed   FlightStatus = "DELAYED"
	Departed  FlightStatus = "DEPARTED"
	Arrived   FlightStatus = "ARRIVED"
	Cancelled FlightStatus = "CANCELLED"
)

type Flight struct {
	FlightNumber  string
	Source        string
	Destination   string
	Status        FlightStatus
	Seats         []Seat
	ArrivalTime   time.Time
	DepartureTime time.Time
	mu            sync.RWMutex
}
