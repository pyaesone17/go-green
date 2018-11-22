package car

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"github.com/pyaesone17/gogreen/app/core"
	"github.com/pyaesone17/gogreen/app/datastore"
	"github.com/pyaesone17/gogreen/app/models"
	"github.com/spf13/viper"
)

// Controller is responsible
type Controller struct {
	config     *viper.Viper
	repository datastore.CarRepository
	core.Controller
}

// Form will bind from JSON
type Form struct {
	Name   string  `form:"name" json:"name" binding:"required"`
	Number string  `form:"number" json:"number" binding:"required"`
	Price  float64 `form:"price" json:"price" binding:"required"`
}

// NewController is constructor to build the interest controller
func NewController(config *viper.Viper) *Controller {
	return &Controller{
		config:     config,
		repository: datastore.NewCarRepository(viper.GetString("mongodb")),
	}
}

// GetCars will command service to return all cars
func (con *Controller) GetCars(c *gin.Context) {
	c.Writer.Write([]byte("Hello world"))
}

// CreateCar will command service to rent the car
func (con *Controller) CreateCar(c *gin.Context) {

	var payload Form

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nextCarID := uuid.New()

	car := models.NewCar(nextCarID, payload.Name, payload.Number, payload.Price)
	err := con.repository.Store(car)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "successfully created", "car": car})
}
