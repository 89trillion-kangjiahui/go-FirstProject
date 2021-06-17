package util

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func SetRedis(key string, value []byte) error {
	rdb := ExampleClient()
	err := rdb.Set(ctx, key, value, 0).Err()
	return err
}

func GetRedis(key string) (string, error) {
	rdb := ExampleClient()
	val, err := rdb.Get(ctx, key).Result()
	return val, err
}
