package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func (handler Handler) Status(context *gin.Context) {
	type response struct{
		Message string `json:"status"`
	}

	res := response{"pong"}

	context.JSON(http.StatusOK, res)
}
