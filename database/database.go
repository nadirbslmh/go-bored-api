package database

import (
	"fmt"
	"go-bored-api/models"
	"go-bored-api/utils"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetConfig("DB_USERNAME"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	DB.AutoMigrate(&models.User{})

	log.Println("connected to the database")
}
