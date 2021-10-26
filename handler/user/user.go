package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/model"
	"net/http"
)

type Handler struct {
	model.UserRepository //@TODO change to service instead
}

func (handler Handler) GetUserList(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		handler.UserRepository.FindAll(),
	)
}
