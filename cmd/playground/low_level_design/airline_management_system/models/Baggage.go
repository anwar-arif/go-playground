package models

type BaggageSize string

const (
	Large  BaggageSize = "LARGE"
	Medium BaggageSize = "MEDIUM"
	Small  BaggageSize = "SMALL"
)

type Baggage struct {
	ID     string
	Size   BaggageSize
	Weight float64
}
