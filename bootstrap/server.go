package bootstrap

import (
	"go-base/api/controllers"
	"go-base/driver"

	"github.com/gin-gonic/gin"
)

type Controller func() *gin.Engine
type Middleware func() gin.HandlerFunc
type ServerConfig struct {
	Loader      []func()
	Controllers []Controller
	Middlewares []gin.HandlerFunc
	Port        string
	DebugLogger bool
}

// CreateServer creates a new Gin server with the given configuration.
// It takes a ServerConfig parameter that specifies the server's configuration.
// The function returns a *gin.Engine, which is the created Gin server.

func CreateServer(config ServerConfig) *gin.Engine {
	if !config.DebugLogger {
		gin.SetMode(gin.ReleaseMode)
	}
	// Create a new Gin server with default middleware
	driver.Instance = gin.Default()

	// Auto apply global middlewares
	applyMiddlewares(driver.Instance, config.Middlewares)

	// Connect the controllers to the server
	controllers.AppController()

	// Run the server on the specified port
	driver.Instance.Run(config.Port)

	// Return the created Gin server
	return driver.Instance
}

// applyMiddlewares applies a list of middlewares to a gin.Engine.
//
// Parameters:
//   - r: a pointer to a gin.Engine object.
//   - middlewares: a slice of gin.HandlerFunc objects.
//
// Return type: None.

func applyMiddlewares(r *gin.Engine, middlewares []gin.HandlerFunc) {
	for _, middleware := range middlewares {
		r.Use(middleware)
	}
}
