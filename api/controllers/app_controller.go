package controllers

import (
	"github.com/gin-gonic/gin"
)

// AppController is a function that handles the routing for the root URL ("/").
//
// It takes in a *gin.RouterGroup parameter named "r" which represents the router group that the function will be added to.
//
// There are no return types specified for this function.
func AppController(r *gin.RouterGroup) {
	appController := r.Group("/")
	appController.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
}
