package bootstrap

import (
	"go-base/api/middlewares"
	"go-base/configs"
	env "go-base/utils"

	"github.com/gin-gonic/gin"
)

// App initializes the application.
//
// It initializes the environment, connects to the database, and creates a server.
func App() *gin.Engine {
	// Initialize the environment
	env.Init()

	// Connect to the database
	DatabaseConnect()
	// Define the server configuration
	serverConfig := ServerConfig{

		Controllers: []Controller{
			//   Controller(controllers.AppController()),
		},

		Middlewares: []gin.HandlerFunc{
			middlewares.Cors(),
			middlewares.Custom(),
		},

		Port: configs.Port,

		DebugLogger: true,
	}

	// Create the server
	return CreateServer(serverConfig)
}
