package router

import (
	"go-example/internal/interfaces/api/middleware"
	"go-example/internal/registry"

	"github.com/gin-gonic/gin"
)

func Root(route *gin.Engine, regi registry.Registry) {
	r := route.Group("/")
	p := r.Group("/private")
	p.Use(middleware.Authorization)
	{
		h := regi.NewUserHandler()
		us := p.Group("/users")
		us.POST("", h.PostUser)
		u := us.Group("/:id")
		u.GET("", h.GetUser)
		u.PATCH("", h.PatchUser)
		u.DELETE("", h.DeleteUser)
	}
}
