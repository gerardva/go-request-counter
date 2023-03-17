package main

import (
	"github.com/gerardva/go-counter-api/config"
	"github.com/gerardva/go-counter-api/controller"
	"github.com/gerardva/go-counter-api/counter"
	"github.com/gerardva/go-counter-api/middleware"
	"github.com/gerardva/go-counter-api/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	initServer()
}

func initServer() {
	conf := config.GetConfig()
	redisClient := redis.NewCacheClient(conf.CacheHost, conf.CachePassword)
	requestCounter := counter.NewRequestCounter(redisClient)

	r := gin.Default()
	r.Use(middleware.CountMiddleware(requestCounter))
	health := new(controller.HealthController)
	r.GET("/health", health.Check)

	count := controller.NewCountController(requestCounter)
	r.GET("/", count.Check)

	r.Run()
}
