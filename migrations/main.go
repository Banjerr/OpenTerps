package migrations

import (
	"log"
	"net/http"
	"os"

	"openterps/dbconnector"
	"openterps/models"

	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

func RunMigration(c *gin.Context) {
	db := dbconnector.ConnectDatabase(os.Getenv("Host"), os.Getenv("Port"), os.Getenv("User"), os.Getenv("Password"), os.Getenv("Name"))

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "initial-migration",
			Migrate: func(tx *gorm.DB) error {
				db.AutoMigrate(&models.Terpene{})
				db.AutoMigrate(&models.Category{})
				db.AutoMigrate(&models.Effect{})
				db.AutoMigrate(&models.Strain{})
				db.AutoMigrate(&models.Smell{})
				db.AutoMigrate(&models.Taste{})
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				tx.Migrator().DropTable("terpenes")
				tx.Migrator().DropTable("categories")
				tx.Migrator().DropTable("effects")
				tx.Migrator().DropTable("strains")
				tx.Migrator().DropTable("smells")
				tx.Migrator().DropTable("tastes")
				tx.Migrator().DropTable("terpene_categories")
				tx.Migrator().DropTable("terpene_effects")
				tx.Migrator().DropTable("terpene_strains")
				tx.Migrator().DropTable("terpene_smells")
				tx.Migrator().DropTable("terpene_tastes")
				return nil
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"data": "Migration ran successfully"})
}
