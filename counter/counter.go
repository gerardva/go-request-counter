package counter

import (
	"fmt"
	"github.com/gerardva/go-counter-api/config"
	"github.com/gerardva/go-counter-api/redis"
)

type RequestCounter struct {
	cache redis.Cache
}

type Counter interface {
	Increment()
	GetTotal() int64
	GetInstanceCount() int64
}

func NewRequestCounter(cache redis.Cache) *RequestCounter {
	return &RequestCounter{
		cache: cache,
	}
}

func (c *RequestCounter) Increment() {
	var configs = config.GetConfig()
	c.cache.Increment(configs.TotalCountKey)
	c.cache.Increment(getInstanceKey())
}

func (c *RequestCounter) GetTotal() int64 {
	var configs = config.GetConfig()
	count, err := c.cache.GetInt(configs.TotalCountKey)
	if err != nil {
		count = 0
	}

	return count
}

func (c *RequestCounter) GetInstanceCount() int64 {
	count, err := c.cache.GetInt(getInstanceKey())
	if err != nil {
		count = 0
	}

	return count
}

func getInstanceKey() string {
	configs := config.GetConfig()
	return fmt.Sprintf("%s:%s", configs.CountKey, configs.Hostname)
}
