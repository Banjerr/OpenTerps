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
				db.AutoMigrate(&models.Benefit{})
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
		{
			ID: "effects-to-benefits-rename",
			Migrate: func(tx *gorm.DB) error {
				type TerpeneBenefit struct {
					TerpeneID int `json:"terpene_id"`
					BenefitID int `json:"benefit_id"`
				}
				db.Migrator().RenameTable("effects", "benefits")
				db.Migrator().RenameTable("terpene_effects", "terpene_benefits")
				db.Migrator().RenameColumn(&TerpeneBenefit{}, "effect_id", "benefit_id")
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				type TerpeneEffect struct {
					TerpeneID int `json:"terpene_id"`
					EffectID  int `json:"effect_id"`
				}
				db.Migrator().RenameTable("benefits", "effects")
				db.Migrator().RenameTable("terpene_benefits", "terpene_effects")
				db.Migrator().RenameColumn(&TerpeneEffect{}, "benefit_id", "effect_id")
				return nil
			},
		},
		{
			ID: "updating-strains-model",
			Migrate: func(tx *gorm.DB) error {
				db.Migrator().AutoMigrate(&models.Strain{})
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				db.Migrator().AutoMigrate(&models.Strain{})
				return nil
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"data": "Migration ran successfully"})
}
