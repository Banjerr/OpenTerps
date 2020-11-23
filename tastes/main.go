package tastes

import (
	"net/http"

	"openterps/dbconnector"
	"openterps/models"
	"openterps/simpleauth"

	"github.com/gin-gonic/gin"
)

// GET /tastes
// Get all tastes
func GetTastes(c *gin.Context) {
	var tastes []models.Taste
	dbconnector.DB.Find(&tastes)

	c.JSON(http.StatusOK, gin.H{"data": tastes})
}

type tastesInput struct {
	Name string `json:"name" binding:"required"`
}

// POST /tastes
// Create a tastes
func CreateTastes(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		// Validate input
		var input tastesInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create tastes
		tastes := models.Taste{
			Name: input.Name,
		}
		dbconnector.DB.Create(&tastes)

		c.JSON(http.StatusOK, gin.H{"data": tastes})
	}
}

// PATCH /tastes/:id
// Update a tastes
func UpdateTastes(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var tastes models.Taste
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&tastes).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Taste not found!"})
			return
		}

		// Validate input
		var input tastesInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dbconnector.DB.Model(&tastes).Updates(input)

		c.JSON(http.StatusOK, gin.H{"data": tastes})
	}
}

// DELETE /tastes/:id
// Delete a tastes
func DeleteTastes(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var tastes models.Taste
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&tastes).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tastes not found!"})
			return
		}

		dbconnector.DB.Delete(&tastes)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
