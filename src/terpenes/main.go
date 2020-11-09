package terpenes

import (
	"net/http"

	"openterps/dbconnector"
	"openterps/models"

	"github.com/gin-gonic/gin"
)

// GET /terpenes
// Get all terpenes
func GetTerpenes(c *gin.Context) {
	var terpenes []models.Terpene
	dbconnector.DB.Find(&terpenes)

	c.JSON(http.StatusOK, gin.H{"data": terpenes})
}
