package datastore

import (
	"math/rand"
	"time"

	"github.com/pyaesone17/gogreen/app/models"
	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CarRepository is the interface to the conversations database
type CarRepository interface {
	Find(id string) (*models.Car, error)
	Store(car *models.Car) error
}

type carRepository struct {
	db *mgo.Session
}

// NewCarRepository creates a new instance of PostDB
func NewCarRepository(mongoDB string) CarRepository {
	db, err := mgo.DialWithTimeout(mongoDB, 60*time.Second)
	if err != nil {
		panic(err)
	}
	db.SetMode(mgo.Eventual, true)
	return &carRepository{
		db: db,
	}
}

func (carRepo *carRepository) Find(id string) (*models.Car, error) {
	type M struct {
		ID  string      `bson:"_id"`
		Car *models.Car `bson:"car"`
	}

	result := &M{}

	collection := carRepo.db.DB(DBName).C(CarCollection)
	err := collection.Find(bson.M{"_id": id}).One(&result)

	if err != nil {
		return nil, err
	}

	return result.Car, nil
}

func (carRepo *carRepository) Store(car *models.Car) error {
	collection := carRepo.db.DB(DBName).C(CarCollection)

	err := collection.Insert(bson.M{
		"_id": car.ID,
		"car": car,
	})

	if err == mgo.ErrNotFound {
		return err
	}
	return nil
}
