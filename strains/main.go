package strains

import (
	"net/http"

	"openterps/dbconnector"
	"openterps/models"
	"openterps/simpleauth"

	"github.com/gin-gonic/gin"
)

// GET /strains
// Get all strains
func GetStrains(c *gin.Context) {
	var strains []models.Strain
	dbconnector.DB.Find(&strains)

	c.JSON(http.StatusOK, gin.H{"data": strains})
}

// GET /strains/:id
// Get a strain by it's ID
func GetStrain(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var strains []models.Strain
	dbconnector.DB.Where("id = ?", id).First(&strains)

	c.JSON(http.StatusOK, gin.H{"data": strains})
}

type strainsInput struct {
	Name string `json:"name" binding:"required"`
}

// POST /strains
// Create a strains
func CreateStrains(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		// Validate input
		var input strainsInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create strains
		strains := models.Strain{
			Name: input.Name,
		}
		dbconnector.DB.Create(&strains)

		c.JSON(http.StatusOK, gin.H{"data": strains})
	}
}

// PATCH /strains/:id
// Update a strains
func UpdateStrains(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var strains models.Strain
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&strains).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Strain not found!"})
			return
		}

		// Validate input
		var input strainsInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dbconnector.DB.Model(&strains).Updates(input)

		c.JSON(http.StatusOK, gin.H{"data": strains})
	}
}

// DELETE /strains/:id
// Delete a strains
func DeleteStrains(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var strains models.Strain
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&strains).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Strains not found!"})
			return
		}

		dbconnector.DB.Delete(&strains)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
