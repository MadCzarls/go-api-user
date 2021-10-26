package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type DataSource struct {
	*redis.Client
}

func (ds *DataSource) Close() error {
	if err := ds.Client.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}

func NewDataSource() *DataSource {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //@TODO move options to env
		Password: "",
		DB:       0,
	})

	return &DataSource{Client: rdb}
}
