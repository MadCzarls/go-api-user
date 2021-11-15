package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mad-czarls/go-api-user/config"
)

type DataSource struct {
	*redis.Client
}

func NewDataSource(cfg config.Config) *DataSource {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	return &DataSource{Client: rdb}
}

func (ds *DataSource) Close() error {
	if err := ds.Client.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}
