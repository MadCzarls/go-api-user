package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/config"
	"github.com/mad-czarls/go-api-user/handler"
	"github.com/mad-czarls/go-api-user/middleware"
)

func SetUpRouter(cfg config.Config, handler handler.Handler) *gin.Engine {
	//@TODO swagger documentation https://github.com/swaggo/gin-swagger

	router := gin.Default()
	router.Use(
		sessions.Sessions(
			cfg.SessionCookieName,
			sessions.NewCookieStore(nil), //"SESSION_HASH_KEY"
		),
	)

	pingGroup := router.Group("/ping")
	{
		pingGroup.GET("", handler.Status)
	}

	api := router.Group("/api")
	{
		userGroup := api.Group("/user")
		{
			userGroup.GET("", handler.GetUserList)
			userGroup.GET("/:id", handler.GetUser)
			userGroup.POST("", handler.Create)
			userGroup.PUT("/:id", handler.Update)
		}
	}

	loginGroup := router.Group("/login")
	{
		loginGroup.POST("", handler.Login)
	}

	profileGroup := router.Group("/profile")
	profileGroup.Use(middleware.AuthMiddleware)
	{
		profileGroup.GET("/me", handler.PersonalInfo)
	}

	logoutGroup := router.Group("/logout")
	{
		logoutGroup.GET("", handler.Logout)
	}

	return router
}
