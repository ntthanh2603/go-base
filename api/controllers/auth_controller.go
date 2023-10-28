package controllers

import (
	"gin-base/api/validation"
	"gin-base/services"

	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func AuthController(r *gin.RouterGroup) {
	authController := r.Group("/auth")
	authService := services.AuthService()
	authController.POST("/login", validation.Body(&LoginDto{}), authService.Login) // Use validation.Body
}
