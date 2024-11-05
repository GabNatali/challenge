package auth

import "github.com/gin-gonic/gin"

func AddAuthRouter(r *gin.Engine, s AuthService) {
	authService := NewAuthHandler(s)

	authGroup := r.Group("/auth")
	authGroup.POST("/login", authService.Login)
}
