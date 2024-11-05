package user

import "github.com/gin-gonic/gin"

func AddUserRouter(r gin.IRouter, usecase UserUseCase) {

	handler := NewUserHandler(usecase)

	r.POST("/users", handler.AddUser)
}
