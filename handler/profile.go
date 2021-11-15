package handler

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/middleware"
)

func (handler *Handler) PersonalInfo(context *gin.Context) {
	session := sessions.Default(context)
	userId := session.Get(middleware.UserKey).(string)

	result, err := handler.UserRepository.FindById(userId)

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"your_id":       result.Id,
			"your_age":      result.Age,
			"your_username": result.Username,
		})
}
