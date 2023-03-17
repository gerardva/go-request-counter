package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type CacheClient struct {
	rdb *redis.Client
}

type Cache interface {
	Increment(key string)
	GetInt(key string) (val int64, err error)
}

func NewCacheClient(host string, password string) *CacheClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})

	var ctx = context.Background()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	return &CacheClient{
		rdb: rdb,
	}
}

func (c *CacheClient) Increment(key string) {
	var ctx = context.Background()
	c.rdb.Incr(ctx, key)
}

func (c *CacheClient) GetInt(key string) (val int64, err error) {
	var ctx = context.Background()
	out, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return
	}

	val, err = strconv.ParseInt(out, 10, 0)
	return
}
