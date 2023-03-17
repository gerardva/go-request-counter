package server

import (
	"github.com/gerardva/go-counter-api/controllers"
	"github.com/gerardva/go-counter-api/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.Use(middleware.CountMiddleware())
	health := new(controllers.HealthController)
	r.GET("/health", health.Check)

	count := new(controllers.CountController)
	r.GET("/", count.Check)

	r.Run()
}
