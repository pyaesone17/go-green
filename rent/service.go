package rent

import (
	"github.com/pborman/uuid"
	"github.com/pyaesone17/gogreen/app/datastore"
	"github.com/pyaesone17/gogreen/app/models"
	"github.com/spf13/viper"
)

// Service will hold domain logic
type Service interface {
	RentCar(carID string, numberOfDays int) (*models.Rent, error)
}

type service struct {
	repository    datastore.RentRepository
	carRepository datastore.CarRepository
}

// NewRentService creates a new instance of Service
func NewRentService(config *viper.Viper) Service {
	return &service{
		carRepository: datastore.NewCarRepository(config.GetString("mongodb")),
		repository:    datastore.NewRentRepository(config.GetString("mongodb")),
	}
}

// RentCar will hold the logic of renting car
func (rentService *service) RentCar(carID string, numberOfDays int) (*models.Rent, error) {
	car, err := rentService.carRepository.Find(carID)
	if err != nil {
		return nil, err
	}

	id := uuid.New()
	// Calcuate the cost based on number
	cost := car.Price * float64(numberOfDays)
	rent := models.NewRent(id, carID, numberOfDays, cost)
	rentService.repository.Store(rent)
	rent.Car = car

	return rent, nil
}
