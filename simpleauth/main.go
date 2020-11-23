package simpleauth

import (
	"net/http"
	"openterps/models"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateRequest(c *gin.Context) bool {
	// validate auth header
	var authHeader models.AuthHeader
	if err := c.ShouldBindHeader(&authHeader); err != nil {
		c.JSON(200, err)
		return false
	}

	if authHeader.APIUser != os.Getenv("APIUser") || authHeader.APIKey != os.Getenv("APIKey") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must be authenticated to do that."})
		return false
	}

	return true
}
