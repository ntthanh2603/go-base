package bootstrap

import "github.com/gin-gonic/gin"

type Controller func(api *gin.RouterGroup)

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
func ConnectControllers(r *gin.Engine, basePath string, ctrls []Controller) {
	api := r.Group(basePath)

	for _, ctrl := range ctrls {
		ctrl(api)
	}
}
