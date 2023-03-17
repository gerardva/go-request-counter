package server

import (
	"github.com/gerardva/go-counter-api/controller"
	"github.com/gerardva/go-counter-api/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.Use(middleware.CountMiddleware())
	health := new(controller.HealthController)
	r.GET("/health", health.Check)

	count := new(controller.CountController)
	r.GET("/", count.Check)

	r.Run()
}
