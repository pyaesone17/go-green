package models

// Car is the root model for blog domain
type Car struct {
	ID        string  `bson:"id" json:"id"`
	Name      string  `json:"name"`
	CarNumber string  `json:"car_number"`
	Price     float64 `json:"price"`
}

// NewCar is the factory for aggregate root
func NewCar(id, name, carNumber string, price float64) *Car {
	return &Car{
		id,
		name,
		carNumber,
		price,
	}
}
