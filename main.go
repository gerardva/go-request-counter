package main

import (
	"github.com/gerardva/go-counter-api/config"
	"github.com/gerardva/go-counter-api/redis"
	"github.com/gerardva/go-counter-api/server"
)

func main() {
	config.Init()
	redis.Init()
	server.Init()
}
