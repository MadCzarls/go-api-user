package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) Status(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "pong",
		},
	)
}
