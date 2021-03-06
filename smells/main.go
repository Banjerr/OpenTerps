package smells

import (
	"net/http"

	"openterps/dbconnector"
	"openterps/models"
	"openterps/simpleauth"

	"github.com/gin-gonic/gin"
)

// GET /smells
// Get all smells
func GetSmells(c *gin.Context) {
	var smells []models.Smell
	dbconnector.DB.Find(&smells)

	c.JSON(http.StatusOK, gin.H{"data": smells})
}

// GET /smells/:id
// Get a smells by it's id
func GetSmell(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var smells []models.Smell
	dbconnector.DB.Where("id = ?", id).First(&smells)

	c.JSON(http.StatusOK, gin.H{"data": smells})
}

type smellsInput struct {
	Name string `json:"name" binding:"required"`
}

// POST /smells
// Create a smells
func CreateSmells(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		// Validate input
		var input smellsInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create smells
		smells := models.Smell{
			Name: input.Name,
		}
		dbconnector.DB.Create(&smells)

		c.JSON(http.StatusOK, gin.H{"data": smells})
	}
}

// PATCH /smells/:id
// Update a smells
func UpdateSmells(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var smells models.Smell
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&smells).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Smell not found!"})
			return
		}

		// Validate input
		var input smellsInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dbconnector.DB.Model(&smells).Updates(input)

		c.JSON(http.StatusOK, gin.H{"data": smells})
	}
}

// DELETE /smells/:id
// Delete a smells
func DeleteSmells(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var smells models.Smell
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&smells).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Smells not found!"})
			return
		}

		dbconnector.DB.Delete(&smells)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
