package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func (handler Handler) Status(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}
