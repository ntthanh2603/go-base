package services

import "github.com/gin-gonic/gin"

type AuthServiceType struct {
}

// AuthService is a function that returns an instance of AuthServiceType.
//
// This function does not take any parameters.
// It returns a pointer to an instance of AuthServiceType.
func AuthService() *AuthServiceType {
	return &AuthServiceType{}
}

// Login handles the login functionality.
//
// It takes a pointer to the AuthServiceType struct and a pointer to the gin.Context struct as parameters.
// It does not return anything.
func (authService *AuthServiceType) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login handled",
	})
}
