package bootstrap

import (
	"gin-base/configs"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Controllers []Controller
	Middlewares []gin.HandlerFunc
	Port        string
}

// CreateServer creates a new Gin server with the given configuration.
// It takes a ServerConfig parameter that specifies the server's configuration.
// The function returns a *gin.Engine, which is the created Gin server.
func CreateServer(config ServerConfig) *gin.Engine {
	// Create a new Gin server with default middleware
	r := gin.Default()

	// Connect the controllers to the server
	ConnectControllers(r, configs.BasePath, config.Controllers)

	// Run the server on the specified port
	r.Run(config.Port)

	// Return the created Gin server
	return r
}
