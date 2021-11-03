package logout

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/middleware"
	"net/http"
)

type Handler struct {
}

func (handler Handler) Logout(context *gin.Context) {
	session := sessions.Default(context)
	userId := session.Get(middleware.UserKey)

	if userId == nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid session token",
			})
		return
	}

	session.Delete(middleware.UserKey)
	if err := session.Save(); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "Failed to save session",
			})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Successfully logged out",
		})
}
