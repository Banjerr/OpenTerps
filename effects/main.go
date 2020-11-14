package effects

import (
	"net/http"

	"openterps/dbconnector"
	"openterps/models"

	"github.com/gin-gonic/gin"
)

// GET /effects
// Get all effects
func GetEffects(c *gin.Context) {
	var effects []models.Effect
	dbconnector.DB.Find(&effects)

	c.JSON(http.StatusOK, gin.H{"data": effects})
}

type effectInput struct {
	Name string `json:"name" binding:"required"`
}

// POST /effect
// Create a effect
func CreateEffect(c *gin.Context) {
	// Validate input
	var input effectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create effect
	effect := models.Effect{
		Name: input.Name,
	}
	dbconnector.DB.Create(&effect)

	c.JSON(http.StatusOK, gin.H{"data": effect})
}

// PATCH /effects/:id
// Update a effect
func UpdateEffect(c *gin.Context) {
	var effect models.Effect
	if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&effect).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Effect not found!"})
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

// DELETE /effects/:id
// Delete a effects
func DeleteEffect(c *gin.Context) {
	var effect models.Effect
	if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&effect).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Effect not found!"})
		return
	}

	dbconnector.DB.Delete(&effect)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
