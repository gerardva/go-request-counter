package main

import (
	"github.com/gerardva/go-counter-api/cache"
	"github.com/gerardva/go-counter-api/config"
	"github.com/gerardva/go-counter-api/server"
)

func main() {
	config.Init()
	cache.Init()
	server.Init()
}
