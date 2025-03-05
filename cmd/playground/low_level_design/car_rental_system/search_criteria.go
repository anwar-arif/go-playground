package car_rental_system

import "time"

type SearchCriteria struct {
	Make      string
	Model     string
	MinPrice  float64
	MaxPrice  float64
	StartDate time.Time
	EndDate   time.Time
}

func NewSearchCriteria(make, model string, minPrice, maxPrice float64, start, end time.Time) *SearchCriteria {
	return &SearchCriteria{
		Make:      make,
		Model:     model,
		MinPrice:  minPrice,
		MaxPrice:  maxPrice,
		StartDate: start,
		EndDate:   end,
	}
}
