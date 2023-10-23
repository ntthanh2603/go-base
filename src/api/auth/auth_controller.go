package auth

import (
	"github.com/gin-gonic/gin"
)

// AuthController is a function that handles the authentication routes.
//
// It takes in a gin.RouterGroup as a parameter.
// There is no return type for this function.
func AuthController(r *gin.RouterGroup) {
	authController := r.Group("/auth")

	authController.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login handled",
		})
	})
}
