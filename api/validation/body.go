package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Body(dto interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		switch contentType {
		case "application/json":
			if err := c.ShouldBindJSON(dto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		case "application/x-www-form-urlencoded":
			if err := c.ShouldBind(dto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported content type"})
			c.Abort()
			return
		}
	}
}
