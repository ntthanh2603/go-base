package main

import (
	"gin-base/src/api/auth"
	"gin-base/src/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	createServer(
		// Pass controllers
		[]routers.Controller{
			auth.AuthController,
		},

		// Pass middlewares
		[]gin.HandlerFunc{},

		// Custom port number
		":3000")
}

// createServer creates a Gin server with the specified controllers, middlewares, and port.
//
// Parameters:
// - controllers: a slice of routers.Controller objects that define the API endpoints.
// - middlewares: a slice of gin.HandlerFunc objects that define the middleware functions.
// - port: a string representing the port number on which the server will listen.
//
// Returns:
// - a pointer to a gin.Engine object representing the Gin server.
func createServer(controllers []routers.Controller, middlewares []gin.HandlerFunc, port string) *gin.Engine {
	r := gin.Default()
	routers.ConnectController(r, "/api", controllers)
	r.Run(port)
	return r
}
