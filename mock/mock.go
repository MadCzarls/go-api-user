package mock

import (
	"github.com/mad-czarls/go-api-user/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u UserRepositoryMock) FindById(id string) (*model.User, error) {
	//transparent method - return the same parameters that will be passed during mocking this method
	args := u.Called()
	user := args.Get(0)

	if user == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (u UserRepositoryMock) FindAll() ([]model.User, error) {
	//transparent method - return the same parameters that will be passed during mocking this method
	args := u.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (u UserRepositoryMock) Create(user *model.User) (*string, error) {
	args := u.Called()
	id := args.Get(0)

	if id == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*string), args.Error(1)
}

func (u UserRepositoryMock) Update(id string, user *model.User) error {
	args := u.Called()

	err := args.Get(0)

	if err == nil {
		return nil
	}

	return args.Error(0)
}