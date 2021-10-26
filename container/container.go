package container

import (
	"github.com/mad-czarls/go-api-user/datasource/redis"
	"github.com/mad-czarls/go-api-user/model"
	redisRepository "github.com/mad-czarls/go-api-user/repository/redis"
)

func GetRedisDataSource() *redis.DataSource {
	return redis.NewDataSource()
}

func GetUserRepository() model.UserRepository {
	return redisRepository.NewUserRepository(GetRedisDataSource())
}
