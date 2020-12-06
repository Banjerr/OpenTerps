package benefits

import (
	"net/http"

	"openterps/dbconnector"
	"openterps/models"
	"openterps/simpleauth"

	"github.com/gin-gonic/gin"
)

// GET /benefits
// Get all benefits
func GetBenefits(c *gin.Context) {
	var benefits []models.Benefit
	dbconnector.DB.Find(&benefits)

	c.JSON(http.StatusOK, gin.H{"data": benefits})
}

// GET /benefits/:id
// Get a benefit by it's ID
func GetBenefit(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var benefits []models.Benefit
	dbconnector.DB.Where("id = ?", id).First(&benefits)

	c.JSON(http.StatusOK, gin.H{"data": benefits})
}

type effectInput struct {
	Name string `json:"name" binding:"required"`
}

// POST /effect
// Create a effect
func CreateBenefit(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		// Validate input
		var input effectInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create effect
		effect := models.Benefit{
			Name: input.Name,
		}
		dbconnector.DB.Create(&effect)

		c.JSON(http.StatusOK, gin.H{"data": effect})
	}
}

// PATCH /benefits/:id
// Update a effect
func UpdateBenefit(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var effect models.Benefit
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&effect).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Benefit not found!"})
			return
		}

		// Validate input
		var input effectInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		dbconnector.DB.Model(&effect).Updates(input)

		c.JSON(http.StatusOK, gin.H{"data": effect})
	}
}

// DELETE /benefits/:id
// Delete a benefits
func DeleteBenefit(c *gin.Context) {
	userIsLoggedIn := simpleauth.ValidateRequest(c)

	if userIsLoggedIn {
		var effect models.Benefit
		if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&effect).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Benefit not found!"})
			return
		}

		dbconnector.DB.Delete(&effect)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
