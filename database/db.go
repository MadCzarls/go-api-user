package database

import (
	"github.com/go-redis/redis/v8"
)

func Client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //@TODO move options to env
		Password: "",
		DB:       0,
	})

	return rdb
}
