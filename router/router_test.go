package router

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/handler"
	"github.com/mad-czarls/go-api-user/middleware"
	"github.com/mad-czarls/go-api-user/mock"
	"github.com/mad-czarls/go-api-user/model"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestSetUpRouter_Profile_ThrowsErrorIfUserUnauthorized(t *testing.T) {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("session_hash_key"))

	router.Use(
		sessions.Sessions(
			"session_cookie_name",
			store,
		),
	)

	userRepository := new(mock.UserRepositoryMock)
	profileHandler := handler.ProfileHandler{UserRepository: userRepository}
	profileGroup := router.Group("/profile")
	profileGroup.Use(middleware.AuthMiddleware)
	{
		profileGroup.GET("/me", profileHandler.PersonalInfo)
	}

	responseWriter := httptest.NewRecorder()

	request := httptest.NewRequest("GET", "/profile/me", nil)

	router.ServeHTTP(responseWriter, request)

	expectedResponse := "{\"error\":\"unauthorized\"}"

	assert.Equal(t, 401, responseWriter.Code)
	assert.Equal(t, expectedResponse, responseWriter.Body.String())
}

func TestSetUpRouter_Profile_Return200IfUserLoggedIn(t *testing.T) {
	router := gin.Default()
	responseWriter := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(responseWriter)
	store := sessions.NewCookieStore([]byte("session_hash_key"))

	router.Use(
		sessions.Sessions(
			"session_cookie_name",
			store,
		),
	)

	user := model.User{Id: "1", Username: "U1", Age: 1}

	userRepository := new(mock.UserRepositoryMock)
	userRepository.On("FindById").Return(user)

	profileHandler := handler.ProfileHandler{UserRepository: userRepository}
	profileGroup := router.Group("/profile")
	//profileGroup.Use(middleware.AuthMiddleware) without middleware - want to test route directly
	{
		profileGroup.GET("/me", profileHandler.PersonalInfo)
	}

	request := httptest.NewRequest("GET", "/profile/me", nil)

	router.ServeHTTP(responseWriter, request)

	session := sessions.Default(testContext)
	userId := session.Get(middleware.UserKey).(string)

	fmt.Print(userId)

	expectedResponse := "{\"error\":\"unauthorized\"}"

	assert.Equal(t, 200, responseWriter.Code)
	assert.Equal(t, expectedResponse, responseWriter.Body.String())
}
