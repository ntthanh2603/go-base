package auth

import "github.com/gin-gonic/gin"

type AuthServiceType struct {
}

func AuthService() *AuthServiceType {
	return &AuthServiceType{}
}

func (authService *AuthServiceType) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login handled",
	})
}
