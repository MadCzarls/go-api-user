package redis

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mad-czarls/go-api-user/datasource/redis"
	"github.com/mad-czarls/go-api-user/model"
)

type UserRepository struct {
	*redis.DataSource
}

func NewUserRepository(dataSource *redis.DataSource) *UserRepository {
	return &UserRepository{DataSource: dataSource}
}

func (u UserRepository) FindById(id string) (*model.User, error) {
	user := &model.User{}
	if err := u.get(id, user); err != nil {
		return nil, err
	}

	if user.Id == "" {
		return nil, nil
	}

	return user, nil
}

func (u UserRepository) FindAll() ([]model.User, error) {
	//@TODO pagination and optimization in the future
	var result []model.User
	ctx := context.Background()

	data, err := u.HGetAll(ctx, "users").Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return result, nil
		}

		return nil, err
	}

	for _, userData := range data {
		user := model.User{}
		json.Unmarshal([]byte(userData), &user)
		result = append(result, user)
	}

	return result, nil
}

func (u UserRepository) Create(user *model.User) error {
	if err := u.save(user); err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Update() {
	panic("implement me")
}

func (u UserRepository) save(value model.Idier) error {
	id := uuid.NewString()
	value.SetId(id)

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()

	err = u.HSet(ctx, "users", id, data).Err()

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) get(key string, dest interface{}) error {
	ctx := context.Background()

	data, err := u.HGet(ctx, "users", key).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		}

		return err
	}

	return json.Unmarshal([]byte(data), dest)
}
