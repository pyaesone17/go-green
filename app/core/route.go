package core

import (
	"github.com/gin-gonic/gin"
)

// Route describe an HTTP route
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
