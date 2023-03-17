package cache

import (
	"context"
	"github.com/gerardva/go-counter-api/config"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var rdb *redis.Client

func Init() {
	configuration := config.GetConfig()
	connect(configuration)
}

func Increment(key string) {
	var ctx = context.Background()
	rdb.Incr(ctx, key)
}

func GetInt(key string) (val int64, err error) {
	var ctx = context.Background()
	out, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return
	}

	val, err = strconv.ParseInt(out, 10, 0)
	return
}

func connect(config *config.Config) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.CacheHost,
		Password: config.CachePassword,
		DB:       0,
	})

	var ctx = context.Background()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
