package services

import "github.com/gin-gonic/gin"

type AppServiceType struct {
}

func AppService() *AppServiceType {
	return &AppServiceType{}
}

func (AppService *AppServiceType) HelloWorldPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "[POST] Hello World",
	})
}

func (AppService *AppServiceType) HelloWorldGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "[GET] Hello World",
	})
}
