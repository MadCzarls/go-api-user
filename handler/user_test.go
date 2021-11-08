package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

type userRepositoryMock struct {
	mock.Mock
}

func (u userRepositoryMock) FindById(id string) (*model.User, error) {
	//transparent method - return the same parameters that will be passed during mocking this method
	args := u.Called()

	user :=  args.Get(0)

	if user == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.User), args.Error(1)
}

func (u userRepositoryMock) FindAll() ([]model.User, error) {
	//transparent method - return the same parameters that will be passed during mocking this method
	args := u.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (u userRepositoryMock) Create(user *model.User) (*string, error) {
	panic("implement me")
}

func (u userRepositoryMock) Update(id string, user *model.User) error {
	panic("implement me")
}

func TestUserHandler_GetUserList_ReturnsListOfUsersProvidedByRepository(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(responseWriter)

	userRepository := new(userRepositoryMock)
	var users []model.User
	user1 := model.User{Id: "1", Username: "U1", Age: 1}
	user2 := model.User{Id: "2", Username: "U2", Age: 2}

	users = append(users, user1)
	users = append(users, user2)
	userRepository.On("FindAll").Return(users, nil)

	handler := UserHandler{userRepository}
	handler.GetUserList(testContext)

	expectedResponse := "[{\"Id\":\"1\",\"username\":\"U1\",\"age\":1},{\"Id\":\"2\",\"username\":\"U2\",\"age\":2}]"

	assert.Equal(t, 200, responseWriter.Code)
	assert.Equal(t, expectedResponse, responseWriter.Body.String())
}

func TestUserHandler_GetUserList_Returns500IfErrorInRepository(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(responseWriter)

	userRepository := new(userRepositoryMock)
	err := errors.New("error thrown in repository")
	userRepository.On("FindAll").Return([]model.User{}, err)

	handler := UserHandler{userRepository}
	handler.GetUserList(testContext)

	expectedResponse := "{\"error\":\"error thrown in repository\"}"

	assert.Equal(t, 500, responseWriter.Code)
	assert.Equal(t, expectedResponse, responseWriter.Body.String())
}

func TestUserHandler_GetUser_Returns500IfErrorInRepository(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	queryParams := gin.Params{
		gin.Param{
			Key:   "id",
			Value: "123",
		},
	}
	testContext, _ := gin.CreateTestContext(responseWriter)
	testContext.Params = queryParams

	userRepository := new(userRepositoryMock)
	err := errors.New("error thrown in repository")
	userRepository.On("FindById").Return(nil, err)

	handler := UserHandler{userRepository}
	handler.GetUser(testContext)

	expectedResponse := "{\"error\":\"error thrown in repository\"}"

	assert.Equal(t, 500, responseWriter.Code)
	assert.Equal(t, expectedResponse, responseWriter.Body.String())
}

func TestUserHandler_GetUser_Returns404IfUserNotFound(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	queryParams := gin.Params{
		gin.Param{
			Key:   "id",
			Value: "123",
		},
	}
	testContext, _ := gin.CreateTestContext(responseWriter)
	testContext.Params = queryParams

	userRepository := new(userRepositoryMock)
	userRepository.On("FindById").Return(nil, nil)

	handler := UserHandler{userRepository}
	handler.GetUser(testContext)

	expectedResponse := "{}"

	assert.Equal(t, 404, responseWriter.Code)
	assert.Equal(t, expectedResponse, responseWriter.Body.String())
}

func TestUserHandler_GetUser_Returns200IfUserFound(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	queryParams := gin.Params{
		gin.Param{
			Key:   "id",
			Value: "123",
		},
	}

	user := model.User{Id: "123", Username: "U123", Age: 123}

	testContext, _ := gin.CreateTestContext(responseWriter)
	testContext.Params = queryParams

	userRepository := new(userRepositoryMock)
	userRepository.On("FindById").Return(&user, nil)

	handler := UserHandler{userRepository}
	handler.GetUser(testContext)

	expectedResponse := "{\"Id\":\"123\",\"username\":\"U123\",\"age\":123}"

	assert.Equal(t, 200, responseWriter.Code)
	assert.Equal(t, expectedResponse, responseWriter.Body.String())
}
