package counter

import (
	"fmt"
	"github.com/gerardva/go-counter-api/config"
	"github.com/gerardva/go-counter-api/redis"
)

func Increment() {
	var configs = config.GetConfig()
	redis.Increment(configs.TotalCountKey)
	redis.Increment(getInstanceKey())
}

func GetTotal() int64 {
	var configs = config.GetConfig()
	count, err := redis.GetInt(configs.TotalCountKey)
	if err != nil {
		count = 0
	}

	return count
}

func GetInstanceCount() int64 {
	count, err := redis.GetInt(getInstanceKey())
	if err != nil {
		count = 0
	}

	return count
}

func getInstanceKey() string {
	configs := config.GetConfig()
	return fmt.Sprintf("%s:%s", configs.CountKey, configs.Hostname)
}
