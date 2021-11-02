package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/container"
	"github.com/mad-czarls/go-api-user/handler/ping"
	"github.com/mad-czarls/go-api-user/handler/user"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//@TODO swagger documentation https://github.com/swaggo/gin-swagger

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
			//@TODO
			//userGroup.PUT("", userHandler.Update)
		}
	}

	//@TODO
	//loginHandler := login.Handler{}
	//
	//loginGroup := router.Group("/login")
	//{
	//	loginGroup.GET("", loginHandler.Status)
	//}

	//@TODO
	//logoutHandler := logout.Handler{}
	//
	//logoutGroup := router.Group("/logout")
	//{
	//	logoutGroup.GET("", logoutHandler.Status)
	//}

	//var inMemoryDb = make(map[string]string)
	//// Get user value
	//router.GET("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := inMemoryDb[user]
	//	if ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//	}
	//})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := router.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if c.Bind(&json) == nil {
	//		inMemoryDb[user] = json.Value
	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})

	return router
}
