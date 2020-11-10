package terpenes

import (
	"log"
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

type terpeneInput struct {
	Name     string            `json:"name" binding:"required"`
	Category []models.Category `json:"category"`
	Taste    []models.Taste    `json:"taste"`
	Smell    []models.Smell    `json:"smell"`
	Strain   []models.Strain   `json:"strain"`
	Effect   []models.Effect   `json:"effect"`
}

// POST /terpenes
// Create a terpene
func CreateTerpene(c *gin.Context) {
	// Validate input
	var input terpeneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("input is ", input)
	// Create terpene
	terpene := models.Terpene{
		Name:     input.Name,
		Category: input.Category,
		Tastes:   input.Taste,
		Smells:   input.Smell,
		Strains:  input.Strain,
		Effects:  input.Effect,
	}
	dbconnector.DB.Create(&terpene)

	c.JSON(http.StatusOK, gin.H{"data": terpene})
}

// PATCH /terpenes/:id
// Update a terpene
func UpdateTerpene(c *gin.Context) {
	var terpene models.Terpene
	if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&terpene).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Terp not found!"})
		return
	}

	// Validate input
	var input terpeneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbconnector.DB.Model(&terpene).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": terpene})
}
