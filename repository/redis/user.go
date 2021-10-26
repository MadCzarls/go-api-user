package redis

import (
	"github.com/mad-czarls/go-api-user/datasource/redis"
)

type UserRepository struct {
	*redis.DataSource
}

func NewUserRepository(dataSource *redis.DataSource) *UserRepository {
	return &UserRepository{DataSource: dataSource}
}

func (u UserRepository) FindById() {
	panic("implement me")
}

func (u UserRepository) FindAll() []string { //@TODO remove after test
	//@TODO example Redis usage below
	//ctx := context.Background()
	//
	//err := db.Set(ctx, "test_key", "test_value2222", 0).Err()
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := db.Get(ctx, "test_key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("test_key value is: ", val)

	return []string{"User1", "User2"}
}

func (u UserRepository) Create() {
	panic("implement me")
}

func (u UserRepository) Update() {
	panic("implement me")
}
