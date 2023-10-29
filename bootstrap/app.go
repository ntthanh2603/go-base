package bootstrap

import (
	"gin-base/api/controllers"
	"gin-base/configs"
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

	serverConfig := ServerConfig{

		Controllers: []Controller{
			controllers.AuthController,
			controllers.AppController,
		},

		Middlewares: []gin.HandlerFunc{},

		Port: configs.Port,
	}

	CreateServer(serverConfig)
}
