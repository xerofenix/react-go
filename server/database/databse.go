package database

import (
	"os"

	"github.com/xerofenix/react-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCon *gorm.DB

// function for connecting to database
func ConnectDB() {
	dsn := os.Getenv("DB_URL")

	var err error
	DBCon, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Error! connecting Database", err)
	}
	println("Databse connected successfully...")

	DBCon.AutoMigrate(new(models.Blog))
}
