package dbconnector

import (
	"fmt"

	"openterps/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO turn this stuff into env vars

var DB *gorm.DB

// ConnectDatabase - connects to database
// returns gorm DB instance, error
func ConnectDatabase(Host, Port, User, Password, Name string) {
	fmt.Println("Host", Host)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to databse: ", dsn)
	}

	db.AutoMigrate(&models.Terpene{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Effect{})
	db.AutoMigrate(&models.Strain{})
	db.AutoMigrate(&models.Smell{})
	db.AutoMigrate(&models.Taste{})

	DB = db
}
