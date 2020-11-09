module openterps

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-gormigrate/gormigrate/v2 v2.0.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
)

// +heroku install github.com/gin-gonic/gin github.com/gin-gonic/gin github.com/jinzhu/gorm github.com/joho/godotenv gorm.io/driver/postgres gorm.io/gorm
