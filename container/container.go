package container

import (
	"github.com/mad-czarls/go-api-user/datasource/redis"
	"github.com/mad-czarls/go-api-user/model"
	redisRepository "github.com/mad-czarls/go-api-user/repository/redis"
	"github.com/mad-czarls/go-api-user/service"
)

var services = make(map[string]interface{})

func GetRedisDataSource() *redis.DataSource {
	if services["redisDS"] == nil {
		//@TODO error handling
		services["redisDS"] = redis.NewDataSource()
	}
	return services["redisDS"].(*redis.DataSource)
}

func GetRedisUserRepository() model.UserRepository {
	if services["get"] == nil {
		//@TODO error handling

		services["userRepo"] = redisRepository.NewUserRepository(GetRedisDataSource())
	}
	return services["userRepo"].(model.UserRepository)
}

func GetEnvManager() (service.VariableGetter, error) {
	if services["envManager"] == nil {
		s, err := service.NewEnvManager()

		if err != nil {
			return nil, err
		}
		services["envManager"] = s
	}
	return services["envManager"].(service.VariableGetter), nil
}
