package user

import "github.com/gin-gonic/gin"

func AddUserRouter(r *gin.Engine, usecase UserUseCase) {

	handler := NewUserHandler(usecase)

	r.POST("/users", handler.AddUser)
}
