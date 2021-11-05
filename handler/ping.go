package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingHandler struct {
}

func (handler PingHandler) Status(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "pong",
		})
}
