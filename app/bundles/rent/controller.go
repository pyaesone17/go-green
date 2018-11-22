package rent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pyaesone17/gogreen/app/core"
	"github.com/pyaesone17/gogreen/rent"
	"github.com/spf13/viper"
)

// Controller is responsible
type Controller struct {
	config  *viper.Viper
	service rent.Service
	core.Controller
}

// Form will bind from JSON
type Form struct {
	CarID        string `json:"car_id" binding:"required"`
	NumberOfDays int    `json:"number_of_days" binding:"required"`
}

// NewController is constructor to build the interest controller
func NewController(config *viper.Viper) *Controller {
	return &Controller{
		config:  config,
		service: rent.NewRentService(config),
	}
}

// GetRent will command service to rent the car
func (con *Controller) GetRent(c *gin.Context) {
	c.Writer.Write([]byte("Hello world"))
}

// RentCar will command service to rent the car
func (con *Controller) RentCar(c *gin.Context) {

	var payload Form

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rent, err := con.service.RentCar(payload.CarID, payload.NumberOfDays)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "successfully created", "rent": rent})
}
