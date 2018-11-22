package rent

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
			Path:    "/rents",
			Handler: c.GetRent,
		},
		core.Route{
			Method:  "POST",
			Path:    "/rents",
			Handler: c.RentCar,
		},
	}
	return &Bundle{r}
}

func (b *Bundle) GetRoutes() []core.Route {
	return b.routes
}
