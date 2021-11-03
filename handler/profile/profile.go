package profile

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/middleware"
	"github.com/mad-czarls/go-api-user/model"
	"net/http"
)

type Handler struct {
	model.UserRepository //@TODO change to service using this repository instead
}

func (handler Handler) PersonalInfo(context *gin.Context) {
	session := sessions.Default(context)
	userId := session.Get(middleware.UserKey).(string)

	result, err := handler.UserRepository.FindById(userId) //@TODO to be handled

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
