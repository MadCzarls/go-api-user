package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/model"
	"net/http"
)

type Handler struct {
	model.UserRepository //@TODO change to service using this repository instead
}

func (handler Handler) GetUserList(context *gin.Context) {
	results, err := handler.UserRepository.FindAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(
		http.StatusOK,
		results,
	)
}

// Create Example cURL request:
// curl 'http://localhost:8080/api/user' -X POST --data-raw '{"username": "John", "age":44}'
func (handler Handler) Create(context *gin.Context) {
	var requestUser model.User
	if err := context.ShouldBindJSON(&requestUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := handler.UserRepository.Create(requestUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, http.NoBody)
}
