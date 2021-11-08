package container

import (
	"github.com/mad-czarls/go-api-user/datasource/redis"
	"github.com/mad-czarls/go-api-user/model"
	redisRepository "github.com/mad-czarls/go-api-user/repository/redis"
	"github.com/mad-czarls/go-api-user/service"
)

var services = make(map[string]interface{})

func GetEnvManager() service.VariableGetter {
	if services["envManager"] == nil {
		s, err := service.NewEnvManager()

		if err != nil {
			panic(err)
		}
		services["envManager"] = s
	}
	return services["envManager"].(service.VariableGetter)
}

func GetRedisDataSource() *redis.DataSource {
	if services["redisDS"] == nil {
		services["redisDS"] = redis.NewDataSource(GetEnvManager())
	}
	return services["redisDS"].(*redis.DataSource)
}

func GetRedisUserRepository() model.UserRepository {
	if services["get"] == nil {
		services["userRepo"] = redisRepository.NewUserRepository(GetRedisDataSource())
	}
	return services["userRepo"].(model.UserRepository)
}
