package terpenes

import (
	"encoding/json"
	"net/http"

	"openterps/dbconnector"
	"openterps/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /terpenes
// Get all terpenes
func GetTerpenes(c *gin.Context) {
	var terpenes []models.Terpene
	dbconnector.DB.Preload("Smells").Preload("Tastes").Preload("Categories").Preload("Strains").Preload("Effects").Find(&terpenes)

	var terpeneResponse []models.TerpeneResponse
	terpeneJSON, _ := json.Marshal(terpenes)
	_ = json.Unmarshal(terpeneJSON, &terpeneResponse)

	c.JSON(http.StatusOK, gin.H{"data": terpeneResponse})
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

	// Create terpene
	terpene := models.Terpene{
		Name:       input.Name,
		Categories: input.Category,
		Tastes:     input.Taste,
		Smells:     input.Smell,
		Strains:    input.Strain,
		Effects:    input.Effect,
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

	terpene.Categories = append(terpene.Categories, input.Category...)
	terpene.Smells = append(terpene.Smells, input.Smell...)
	terpene.Tastes = append(terpene.Tastes, input.Taste...)
	terpene.Effects = append(terpene.Effects, input.Effect...)
	terpene.Strains = append(terpene.Strains, input.Strain...)

	dbconnector.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&terpene)

	c.JSON(http.StatusOK, gin.H{"data": &terpene})
}

// DELETE /terpenes/:id
// Delete a terpene
func DeleteTerpene(c *gin.Context) {
	var terpene models.Terpene
	if err := dbconnector.DB.Where("id = ?", c.Param("id")).First(&terpene).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Terpene not found!"})
		return
	}

	dbconnector.DB.Delete(&terpene)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
