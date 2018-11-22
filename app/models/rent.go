package models

// Rent is the root model
type Rent struct {
	ID           string  `json:"id"`
	CarID        string  `json:"car_id"`
	NumberOfDays int     `json:"number_of_days"`
	Cost         float64 `json:"cost"`
}

// NewRent is the factory for aggregate root
func NewRent(ID, carID string, numberOfDays int, cost float64) *Rent {
	return &Rent{
		ID,
		carID,
		numberOfDays,
		cost,
	}
}
