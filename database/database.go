package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       dbNo,
	})
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
