package counter

import (
	"fmt"
	"github.com/gerardva/go-counter-api/cache"
	"github.com/gerardva/go-counter-api/config"
)

func Increment() {
	var configs = config.GetConfig()
	cache.Increment(configs.TotalCountKey)
	cache.Increment(getInstanceKey())
}

func GetTotal() int64 {
	var configs = config.GetConfig()
	count, err := cache.GetInt(configs.TotalCountKey)
	if err != nil {
		count = 0
	}

	return count
}

func GetInstanceCount() int64 {
	count, err := cache.GetInt(getInstanceKey())
	if err != nil {
		count = 0
	}

	return count
}

func getInstanceKey() string {
	configs := config.GetConfig()
	return fmt.Sprintf("%s:%s", configs.CountKey, configs.Hostname)
}
