package car

import (
	"github.com/pyaesone17/gogreen/app/core"
	"github.com/spf13/viper"
)

type Bundle struct {
	routes []core.Route
}

func NewBundle(config *viper.Viper) core.Bundle {
	c := NewController(config)
	r := []core.Route{
		core.Route{
			Method:  "GET",
			Path:    "/cars",
			Handler: c.GetCars,
		},
		core.Route{
			Method:  "POST",
			Path:    "/cars",
			Handler: c.CreateCar,
		},
	}
	return &Bundle{r}
}

func (b *Bundle) GetRoutes() []core.Route {
	return b.routes
}
