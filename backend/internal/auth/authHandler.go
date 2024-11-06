package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
	authService AuthService
}

func NewAuthHandler(service AuthService) AuthHandler {
	return &authHandler{
		authService: service,
	}
}

func (a *authHandler) Login(c *gin.Context) {

	var loginRequest LoginUserDto

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := a.authService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User logged in successfully",
		"data":    user,
	})
}
