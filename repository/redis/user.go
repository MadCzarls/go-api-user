package redis

import (
	"context"
	"encoding/json"
	"errors"
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

func (u UserRepository) Create(user *model.User) (*string, error) {
	id := uuid.NewString()
	user.SetId(id)

	if err := u.save(user); err != nil {
		return nil, err
	}
	return &id, nil
}

func (u UserRepository) Update(id string, newUserData *model.User) error {
	currentUser, err := u.FindById(id)

	if err != nil {
		return err
	}

	if currentUser == nil {
		return errors.New("user does not exist")
	}

	newUserData.SetId(currentUser.GetId())

	if err := u.save(newUserData); err != nil {
		return err
	}
	return nil
}

func (u UserRepository) save(value model.Idier) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()

	err = u.HSet(ctx, "users", value.GetId(), data).Err()

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
