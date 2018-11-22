package models

import (
	"encoding/json"
)

// Rent is the root model
type Rent struct {
	ID           string  `json:"id,omitempty"`
	CarID        string  `json:"car_id,omitempty"`
	NumberOfDays int     `json:"number_of_days,omitempty"`
	Cost         float64 `json:"cost,omitempty"`
	Car          *Car    `json:"car,omitempty"`
}

// NewRent is the factory for aggregate root
func NewRent(ID, carID string, numberOfDays int, cost float64) *Rent {
	return &Rent{
		ID:           ID,
		CarID:        carID,
		NumberOfDays: numberOfDays,
		Cost:         cost,
	}
}

// JSON will transform Rent model to readable
func (rent *Rent) JSON() *Rent {
	return &Rent{
		CarID:        rent.CarID,
		NumberOfDays: rent.NumberOfDays,
		Cost:         rent.Cost,
	}
}

// Rules will return the rules
func (rent *Rent) Rules() map[string]interface{} {
	carRules := []interface{}{
		"car", rent.Car.Rules(),
	}

	data := map[string]interface{}{
		"car_id":         "carID|string",
		"number_of_days": "NumberOfDays|string",
		"unique": func(data interface{}) interface{} {
			var rent Rent
			bytes, _ := json.Marshal(data)
			json.Unmarshal(bytes, &rent)
			return rent.ID + "_" + rent.CarID
		},
		"car": carRules,
	}
	return data
}
