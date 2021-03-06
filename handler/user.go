package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mad-czarls/go-api-user/model"
	"net/http"
)

type UserHandler struct {
	model.UserRepository
}

func (handler UserHandler) GetUserList(context *gin.Context) {
	results, err := handler.UserRepository.FindAll()

	if err != nil {
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
		results,
	)
}

func (handler UserHandler) GetUser(context *gin.Context) {
	result, err := handler.UserRepository.FindById(context.Param("id"))

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if result != nil {
		context.JSON(
			http.StatusOK,
			result,
		)
		return
	}

	context.JSON(
		http.StatusNotFound,
		http.NoBody,
	)
}

// Create Example cURL request:
// curl 'http://localhost:8080/api/user' -X POST --data-raw '{"username": "John", "age":44}'
func (handler UserHandler) Create(context *gin.Context) {
	var requestUser model.User
	if err := context.ShouldBindJSON(&requestUser); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}
	id, err := handler.UserRepository.Create(&requestUser)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"id": id,
		},
	)
}

// Update Example cURL request:
// curl 'http://localhost:8080/api/user/306ba65d-a4b8-4ebb-a30b-93526b31b8d9' -X PUT --data-raw '{"username": "John", "age":44}'
func (handler UserHandler) Update(context *gin.Context) {
	var requestUser model.User
	if err := context.ShouldBindJSON(&requestUser); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	if err := handler.UserRepository.Update(context.Param("id"), &requestUser); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		http.NoBody,
	)
}
