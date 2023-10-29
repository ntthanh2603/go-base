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

func CreateServer(config ServerConfig) *gin.Engine {
	r := gin.Default()
	ConnectControllers(r, configs.BasePath, config.Controllers)
	r.Run(config.Port)
	return r
}
