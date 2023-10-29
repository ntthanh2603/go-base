package bootstrap

import (
	"gin-base/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller func(api *gin.RouterGroup)
type Middleware func() gin.HandlerFunc
type ServerConfig struct {
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
	r := gin.Default()

	// Auto apply global middlewares
	applyMiddlewares(r, config.Middlewares)

	// Connect the controllers to the server
	connectControllers(r, configs.BasePath, config.Controllers)

	// Run the server on the specified port
	r.Run(config.Port)

	// Return the created Gin server
	return r
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

// ConnectController connects the given controllers to the provided gin.Engine.
//
// It creates a new gin.RouterGroup under the specified basePath and iterates over
// the list of controllers, calling each controller function with the router group as
// the argument.
//
// Parameters:
// - r: The gin.Engine instance to connect the controllers to.
// - basePath: The base path for the router group.
// - ctrls: A slice of Controller functions to connect.

// connectControllers registers the controllers with the specified base path
func connectControllers(r *gin.Engine, basePath string, ctrls []Controller) {
	// Redirect root path to the base path if it is not empty
	if basePath != "" {
		r.Any("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, basePath)
		})
	}

	// Create a new router group with the base path
	api := r.Group(basePath)

	// Register each controller with the router group
	for _, ctrl := range ctrls {
		ctrl(api)
	}
}
