package main

import (
	"net/http"
	"os"

	"openterps/dbconnector"
	"openterps/migrations"
	"openterps/terpenes"

	"github.com/gin-gonic/gin"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Print("Error loading .env file. Hopefully this isn't local")
	// }

	r := gin.Default()

	dbconnector.ConnectDatabase(os.Getenv("Host"), os.Getenv("Port"), os.Getenv("User"), os.Getenv("Password"), os.Getenv("Name"))

	// GET / - home/info
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to OpenTerps 0.0.1!"})
	})

	// GET /terpenes
	r.GET("/terpenes", terpenes.GetTerpenes)

	// POST /migrations
	r.POST("/migrations", migrations.RunMigration)

	r.Run(":" + os.Getenv("PORT"))
}
