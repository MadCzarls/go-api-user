package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mad-czarls/go-api-user/service"
)

type DataSource struct {
	*redis.Client
	service.VariableGetter
}

func (ds *DataSource) Close() error {
	if err := ds.Client.Close(); err != nil {
		return fmt.Errorf("error closing Redis Client: %w", err)
	}

	return nil
}

func NewDataSource(envManager service.VariableGetter) *DataSource {
	addr := envManager.GetEnvString("REDIS_HOST")
	password := envManager.GetEnvString("REDIS_PASSWORD")
	db := envManager.GetEnvInt("REDIS_DATABASE")

	rdb := redis.NewClient(&redis.Options{
		Addr:     *addr,
		Password: *password,
		DB:       *db,
	})

	return &DataSource{Client: rdb}
}
