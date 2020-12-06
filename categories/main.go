package categories

import (
	"net/http"

	"openterps/dbconnector"
	"openterps/models"

	"github.com/gin-gonic/gin"
)

// GET /categories
// Get all categories
func GetCategories(c *gin.Context) {
	var categories []models.Category
	dbconnector.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GET /categories/:id
// Get a category by its ID
func GetCategory(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var categories []models.Category
	dbconnector.DB.Where("id = ?", id).First(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

type categoriesInput struct {
	Name string `json:"name" binding:"required"`
}

// POST /categories
// Create a categories
func CreateCategories(c *gin.Context) {
	// Validate input
	var input categoriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create categories
	categories := models.Category{
		Name: input.Name,
	}
	dbconnector.DB.Create(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// PATCH /categories/:id
// Update a categories
func UpdateCategories(c *gin.Context) {
	var categories models.Category
	if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!"})
		return
	}

	// Validate input
	var input categoriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbconnector.DB.Model(&categories).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// DELETE /categories/:id
// Delete a categories
func DeleteCategories(c *gin.Context) {
	var categories models.Category
	if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!"})
		return
	}

	dbconnector.DB.Delete(&categories)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
