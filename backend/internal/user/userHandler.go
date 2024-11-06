package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	AddUser(c *gin.Context)
}

type userHandler struct {
	userUseCase UserUseCase
}

func NewUserHandler(cases UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: cases,
	}
}

func (u *userHandler) AddUser(c *gin.Context) {

	var addUserDto AddUserDto

	if err := c.ShouldBindJSON(&addUserDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	userId, err := u.userUseCase.Add(addUserDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User added successfully",
		"data": map[string]interface{}{
			"id": userId,
		},
	})
}
