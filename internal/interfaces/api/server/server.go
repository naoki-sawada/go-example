package server

import (
	"go-example/internal/interfaces/api/router"
	"go-example/internal/registry"

	"github.com/gin-gonic/gin"
)

func Create(r registry.Registry) *gin.Engine {
	app := gin.Default()
	router.Root(app, r)
	return app
}
