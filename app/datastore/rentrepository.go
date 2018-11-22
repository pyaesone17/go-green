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

// RentRepository is the interface to the conversations database
type RentRepository interface {
	Store(rent *models.Rent) error
}

type rentRepository struct {
	db *mgo.Session
}

// NewRentRepository creates a new instance of PostDB
func NewRentRepository(mongoDB string) RentRepository {
	db, err := mgo.DialWithTimeout(mongoDB, 60*time.Second)
	if err != nil {
		panic(err)
	}
	db.SetMode(mgo.Eventual, true)
	return &rentRepository{
		db: db,
	}
}

func (rentRepo *rentRepository) Store(rent *models.Rent) error {

	collection := rentRepo.db.DB(DBName).C(RentCollection)
	err := collection.Insert(bson.M{
		"_id":  rent.ID,
		"rent": rent,
	})

	if err == mgo.ErrNotFound {
		return err
	}
	return nil
}
