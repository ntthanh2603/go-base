package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Body is a Go function that takes in a dto parameter and returns a gin.HandlerFunc.
//
// The function is responsible for validating a JSON request body using the provided dto structure.
// It first attempts to bind the JSON request body to the dto structure using c.ShouldBindJSON.
// If the binding fails, it returns a HTTP 400 Bad Request response with the error message.
// If the binding succeeds, it validates the dto structure using a validator and returns a HTTP 400 Bad Request response with the error message if validation fails.
// If the binding and validation both succeed, the function does not return anything.
func Body(dto interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var validate = validator.New()
		if err := c.ShouldBindJSON(dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if err := validate.Struct(dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
	}
}
