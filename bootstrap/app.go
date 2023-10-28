package bootstrap

import (
	"gin-base/api/controllers"
	env "gin-base/utils"

	"github.com/gin-gonic/gin"
)

// App initializes the application.
//
// It initializes the environment, connects to the database, and creates a server.
// The function
func App() {
	env.Init()
	DatabaseConnect()
	createServer(
		// Pass controllers
		[]Controller{
			controllers.AuthController,
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
func createServer(controllers []Controller, middlewares []gin.HandlerFunc, port string) *gin.Engine {
	r := gin.Default()
	ConnectControllers(r, "/api", controllers)
	r.Run(port)
	return r
}
