package handler

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/middleware"
	"github.com/mad-czarls/go-api-user/model"
	"net/http"
)

type LoginHandler struct {
	model.UserRepository
}

func (handler LoginHandler) Login(context *gin.Context) {
	var authData model.Auth
	if err := context.ShouldBindJSON(&authData); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	result, err := handler.UserRepository.FindById(authData.Id)
	//@TODO implement password handling in the future

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if result == nil {
		context.JSON(
			http.StatusNotFound,
			http.NoBody,
		)
		return
	}

	session := sessions.Default(context)
	session.Set(middleware.UserKey, authData.Id)

	if err = session.Save(); err != nil {
		context.JSON(
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
			"message": "You are logged in!",
		},
	)
}
