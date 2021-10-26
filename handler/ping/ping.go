package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func (handler Handler) Status(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "pong",
		})
}
