package gogreen

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pyaesone17/gogreen/app/bundles/car"
	"github.com/pyaesone17/gogreen/app/bundles/rent"
	"github.com/pyaesone17/gogreen/app/core"
	viper "github.com/spf13/viper"
)

// BlogService holds
type BlogService struct {
	viper *viper.Viper
	srv   *http.Server
}

// NewService is constructor for BlogService struct
func NewService(viper *viper.Viper) *BlogService {
	return &BlogService{
		viper: viper,
	}
}

// ListenAndServe will listen the port
func (s *BlogService) ListenAndServe() {
	r := gin.Default()
	for _, b := range initBundles(s.viper) {
		for _, route := range b.GetRoutes() {
			log.Printf("adding handler for \"%s %s\"", route.Method, route.Path)
			r.Handle(route.Method, route.Path, route.Handler)
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080
	return
}

// Stop will stop running the server
func (s *BlogService) Stop() {
	if s.srv != nil {
		s.srv.Close()
	}
}

func initBundles(viper *viper.Viper) []core.Bundle {
	return []core.Bundle{
		rent.NewBundle(viper),
		car.NewBundle(viper),
	}
}
