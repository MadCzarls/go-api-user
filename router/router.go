package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/container"
	"github.com/mad-czarls/go-api-user/handler/login"
	"github.com/mad-czarls/go-api-user/handler/logout"
	"github.com/mad-czarls/go-api-user/handler/ping"
	"github.com/mad-czarls/go-api-user/handler/profile"
	"github.com/mad-czarls/go-api-user/handler/user"
	"github.com/mad-czarls/go-api-user/middleware"
	"github.com/mad-czarls/go-api-user/session"
)

func SetUpRouter() *gin.Engine {
	//@TODO swagger documentation https://github.com/swaggo/gin-swagger

	router := gin.Default()
	router.Use(sessions.Sessions("app_session", session.SetUpSession())) //@TODO put name in ENV

	pingHandler := ping.Handler{}
	pingGroup := router.Group("/ping")
	{
		pingGroup.GET("", pingHandler.Status)
	}

	api := router.Group("/api")
	{
		userHandler := user.Handler{UserRepository: container.GetUserRepository()}
		userGroup := api.Group("/user")
		{
			userGroup.GET("", userHandler.GetUserList)
			userGroup.GET("/:id", userHandler.GetUser)
			userGroup.POST("", userHandler.Create)
			userGroup.PUT("/:id", userHandler.Update)
		}
	}

	loginHandler := login.Handler{UserRepository: container.GetUserRepository()}
	loginGroup := router.Group("/login")
	{
		loginGroup.POST("", loginHandler.Login)
	}

	profileHandler := profile.Handler{UserRepository: container.GetUserRepository()}
	profileGroup := router.Group("/profile")
	profileGroup.Use(middleware.AuthMiddleware)
	{
		profileGroup.GET("/me", profileHandler.PersonalInfo)
	}

	logoutHandler := logout.Handler{}

	logoutGroup := router.Group("/logout")
	{
		logoutGroup.GET("", logoutHandler.Logout)
	}

	return router
}
