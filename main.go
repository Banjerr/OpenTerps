package main

import (
	"log"
	"net/http"
	"os"

	"openterps/benefits"
	"openterps/categories"
	"openterps/dbconnector"
	"openterps/migrations"
	"openterps/smells"
	"openterps/strains"
	"openterps/tastes"
	"openterps/terpenes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file. Hopefully this isn't local")
	}

	r := gin.Default()

	dbconnector.ConnectDatabase(os.Getenv("Host"), os.Getenv("Port"), os.Getenv("User"), os.Getenv("Password"), os.Getenv("Name"))

	// GET / - home/info
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to OpenTerps 0.0.1!"})
	})

	// GET /terpenes
	r.GET("/terpenes", terpenes.GetTerpenes)

	// POST /terpenes
	r.POST("/terpenes", terpenes.CreateTerpene)

	// PUT /terpenes
	r.PATCH("/terpenes/:id", terpenes.UpdateTerpene)

	// Delete /terpenes
	r.DELETE("/terpenes/:id", terpenes.DeleteTerpene)

	// GET /benefits
	r.GET("/benefits", benefits.GetBenefits)

	// POST /benefits
	r.POST("/benefits", benefits.CreateBenefit)

	// PUT /benefits
	r.PATCH("/benefits/:id", benefits.UpdateBenefit)

	// Delete /benefits
	r.DELETE("/benefits/:id", benefits.DeleteBenefit)

	// GET /tastes
	r.GET("/tastes", tastes.GetTastes)

	// POST /tastes
	r.POST("/tastes", tastes.CreateTastes)

	// PUT /tastes
	r.PATCH("/tastes/:id", tastes.UpdateTastes)

	// Delete /tastes
	r.DELETE("/tastes/:id", tastes.DeleteTastes)

	// GET /smells
	r.GET("/smells", smells.GetSmells)

	// POST /smells
	r.POST("/smells", smells.CreateSmells)

	// PUT /smells
	r.PATCH("/smells/:id", smells.UpdateSmells)

	// Delete /smells
	r.DELETE("/smells/:id", smells.DeleteSmells)

	// GET /strains
	r.GET("/strains", strains.GetStrains)

	// POST /strains
	r.POST("/strains", strains.CreateStrains)

	// PUT /strains
	r.PATCH("/strains/:id", strains.UpdateStrains)

	// Delete /strains
	r.DELETE("/strains/:id", strains.DeleteStrains)

	// GET /categories
	r.GET("/categories", categories.GetCategories)

	// POST /categories
	r.POST("/categories", categories.CreateCategories)

	// PUT /categories
	r.PATCH("/categories/:id", categories.UpdateCategories)

	// Delete /categories
	r.DELETE("/categories/:id", categories.DeleteCategories)

	// POST /migrations
	r.POST("/migrations", migrations.RunMigration)

	r.Run(":" + os.Getenv("PORT"))
}
