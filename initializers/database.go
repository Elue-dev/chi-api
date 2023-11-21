package initializers

import (
	"fmt"
	"os"

	"github.com/elue-dev/gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectAndMigrateDB() {
	 dsn := os.Getenv("DB_DSN")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Could not connect to DB")
	}

	fmt.Println("Connected to Postgres DB")

	if err = DB.AutoMigrate(&models.Post{}); err != nil {
		fmt.Println("DB migration failed", err)
	} 
}

