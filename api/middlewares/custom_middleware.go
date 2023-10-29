package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Custom() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic handle middleware

		//
		c.Next()
		return
	}
}
