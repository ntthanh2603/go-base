package bootstrap

import (
	"gin-base/api/controllers"
	"gin-base/api/middlewares"
	"gin-base/configs"
	env "gin-base/utils"

	"github.com/gin-gonic/gin"
)

// App initializes the application.
//
// It initializes the environment, connects to the database, and creates a server.
func App() {
	// Initialize the environment
	env.Init()

	// Connect to the database
	DatabaseConnect()

	// Define the server configuration
	serverConfig := ServerConfig{

		Controllers: []Controller{
			controllers.AuthController,
			controllers.AppController,
		},

		Middlewares: []gin.HandlerFunc{
			middlewares.Cors(),
			middlewares.Custom(),
		},

		Port: configs.Port,

		DebugLogger: true,
	}

	// Create the server
	CreateServer(serverConfig)
}
