package middleware

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserKey = "user"

func AuthMiddleware(context *gin.Context) {
	session := sessions.Default(context)

	userId := session.Get(UserKey)

	if userId == nil {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "unauthorized",
			},
		)
		return
	}

	context.Next()
}
